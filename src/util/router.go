package util

import (
	"embed"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

var (
	Engine           *gin.Engine
	fs               http.FileSystem
	fileServer       http.Handler
	client           = resty.New()
	Static           embed.FS
	FindUserInfoById func(id int) any
)

func RouterRegister() {
	// Engine.StaticFS("/vue", http.FS(Static))
	Engine.Any("/ui/*filepath", staticFs)
	// 匹配vue中的/v/*链接，跳转至vue入口文件，vue会自动进行路由
	Engine.GET("/ui", getUi)
	// 匹配/链接，重定向到主页
	Engine.GET("/", firstPage)
	Engine.PUT("/:context/:libName/*filePath", put)
	Engine.GET("/:context/:libName/*filePath", get)
	Engine.HEAD("/:context/:libName/*filePath", get)
}
func staticFs(c *gin.Context) {
	staticServer := http.FileServer(http.FS(Static))
	staticServer.ServeHTTP(c.Writer, c.Request)
}
func getUi(c *gin.Context) {
	c.Request.URL.Path = "/ui/index.html"
	Engine.HandleContext(c)
}
func firstPage(c *gin.Context) {
	c.Redirect(http.StatusFound, "/ui/")
}

func get(c *gin.Context) {
	repository, err := checkAndGetRepository(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	if repository.Mode&4 != 4 {
		c.String(http.StatusForbidden, "repository not support read")
		return
	}

	filePath := c.Param("filePath")
	ext := path.Ext(filePath)
	if ext == "" && !strings.HasSuffix(filePath, "/") {
		c.Redirect(http.StatusMovedPermanently, c.Request.RequestURI+"/")
		return
	}

	localFilePath := path.Join(repository.Target, filePath)

	f, err := fs.Open(localFilePath)
	defer closeFile(f)

	localFilePath = path.Join(config.LocalRepository, localFilePath)
	if err != nil && len(repository.Mirror) > 0 {
		// 尝试从url镜像获取返回
		response := readRemote(repository, filePath)
		if response == nil {
			c.String(http.StatusNotFound, "not found")
		}

		data := response.Body()
		status := response.StatusCode()
		if repository.Cache && status == http.StatusOK {
			// 不缓存metadata
			filePath = strings.ToLower(filePath)
			if !strings.Contains(filePath, "maven-metadata.xml") {
				if err = saveFile(localFilePath, data); err != nil {
					log.Errorf("cache mirror file failed. message: %v", err)
				}
			}
		}

		c.Data(response.StatusCode(), response.Header().Get("Content-Type"), data)
		return
	}

	if generate := c.Query("generate_md5_sha1"); strings.EqualFold(generate, "true") {
		if !checkAuth(c) {
			c.String(http.StatusUnauthorized, "Unauthorised")
			return
		}

		if err = generateHash(localFilePath); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("generate hash failed, message: %v\n", err))
		}
	}

	if repository.Target != repository.Id {
		u := fmt.Sprintf("/%s/%s%s", config.Context, repository.Target, filePath)
		c.Request.URL.RawPath = u
		c.Request.URL.Path = u
	}

	fileServer.ServeHTTP(c.Writer, c.Request)
}

func put(c *gin.Context) {
	if !checkAuth(c) {
		c.String(http.StatusUnauthorized, "Unauthorised")
		return
	}

	length, err1 := strconv.Atoi(c.GetHeader("Content-Length"))
	data, err2 := io.ReadAll(c.Request.Body)
	if err1 != nil || err2 != nil || length <= 0 || length != len(data) {
		log.Errorf("data read failed%v\n%v", err1, err2)
		c.String(http.StatusInternalServerError, "data read failed")
		return
	}

	repository, err := checkAndGetRepository(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	if repository.Mode&2 != 2 {
		c.String(http.StatusForbidden, "repository not support write")
		return
	}

	filePath := c.Param("filePath")
	localFilePath := path.Join(config.LocalRepository, repository.Target, filePath)
	if err = saveFile(localFilePath, data); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("write file failed. message: %v\n", err))
		return
	}

	if generate := c.Query("generate_md5_sha1"); strings.EqualFold(generate, "true") {
		if err = generateHash(localFilePath); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("generate hash failed, message: %v\n", err))
		}
	}

	c.String(http.StatusOK, "OK")
}

func saveFile(localFilePath string, data []byte) error {
	if err := CreateParentIfNotExist(localFilePath); err != nil {
		return err
	}

	if err := os.WriteFile(localFilePath, data, 0755); err != nil {
		return err
	}
	return nil
}

func readRemote(repository *Repository, filePath string) *resty.Response {
	for _, u := range repository.Mirror {
		u = fmt.Sprintf("%s%s", u, filePath)
		if response, err := client.R().Get(u); err != nil {
			log.Errorf("request mirror url '%s' failed", u)
		} else {
			return response
		}
	}
	return nil
}

func generateHash(file string) error {
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		dir, err := os.ReadDir(file)
		if err != nil {
			return err
		}
		for _, info := range dir {
			if err = generateHash(info.Name()); err != nil {
				return err
			}
		}
	}
	ext := path.Ext(file)
	if ext != ".xml" && ext != ".jar" && ext != ".pom" {
		return nil
	}
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	if err = touchFile(file, "md5", bytes); err != nil {
		return err
	}
	if err = touchFile(file, "sha1", bytes); err != nil {
		return err
	}
	return nil
}

func touchFile(file string, hash string, bytes []byte) error {
	hashFile := fmt.Sprintf("%s.%s", file, hash)
	if exist, err := CheckFileExist(hashFile); err != nil {
		return err
	} else if !exist {
		if err = os.WriteFile(hashFile, GetHash(bytes, hash), 0755); err != nil {
			return err
		}
	}
	return nil
}

func closeFile(f http.File) {
	if f != nil {
		_ = f.Close()
	}
}

func checkAndGetRepository(c *gin.Context) (repository *Repository, err error) {
	context := c.Param("context")
	libName := c.Param("libName")
	filePath := c.Param("filePath")

	if context == "" || libName == "" {
		return nil, errors.New("empty repository")
	}

	fullPath := fmt.Sprintf("/%s/%s%s", context, libName, filePath)
	if context != config.Context {
		return nil, errors.New(fmt.Sprintf("not found, url = %s", fullPath))
	}

	// 获取存储库配置
	repository = config.RepositoryStore[libName]
	if repository == nil {
		return nil, errors.New(fmt.Sprintf("repository %s is not actived", libName))
	}

	return repository, nil
}

func checkAuth(c *gin.Context) bool {
	authorization := c.GetHeader("Authorization")
	if !strings.HasPrefix(authorization, "Basic ") {
		return false
	}
	// 校验用户
	authorization = strings.TrimSpace(authorization[6:])
	if config.Auth[authorization] == nil {
		return false
	}
	return true
}

func jwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		curl := c.Request.RequestURI
		if authExcludeRegexp.MatchString(curl) {
			c.Next()
		} else {
			token := c.Request.Header.Get("Authorization")
			if token != "" {
				state, id, act := ValidToken(token)
				log.Infof("state:%s id:%s act:%s", state, id, act)
				if state {
					currentUser := FindUserInfoById(id)
					c.Set("id", id)
					c.Set("act", act)
					c.Set("current", currentUser)
					c.Next()
					return
				}
			}
			c.JSON(http.StatusOK, FAIL(MsgCodeTokenValidFail, nil))
			c.Abort()

		}

	}
}
