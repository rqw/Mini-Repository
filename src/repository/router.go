package repository

import (
	"Mini-Repository/src/util"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RouterRegister() {
	util.Engine.GET("/repository/:id", _getRepository)
	util.Engine.DELETE("/repository/:id", _dropRepository)
	util.Engine.PUT("/repository", _saveRepository)
	util.Engine.POST("/repository", _queryRepository)
	util.Engine.POST("/repository/view/:libName", _viewComponent)
	util.Engine.POST("/repository/del/:libName", _deleteComponent)
	util.Engine.POST("/repository/upload/:libName/*filePath", _uploadComponent)
	util.Engine.PUT("/:context/:libName/*filePath", put)
	util.Engine.GET("/:context/:libName/*filePath", get)
	util.Engine.HEAD("/:context/:libName/*filePath", get)
}
func _getRepository(c *gin.Context) {
	if id, err := util.GetParamId(c); err == nil {
		info := cache[id]
		c.JSON(http.StatusOK, util.SUCCESS(info))
	}
}
func _dropRepository(c *gin.Context) {
	if id, err := util.GetParamId(c); err == nil {
		DelRepository(id)
		c.JSON(http.StatusOK, util.SUCCESS(nil))
	}
}
func _saveRepository(c *gin.Context) {
	if repos, err := util.GetParamJson[Repository](c); err == nil {
		code := SaveRepository(&repos)
		if code == util.MsgCodeSuccess {
			c.JSON(http.StatusOK, util.SUCCESS(nil))
		} else {
			c.JSON(http.StatusOK, util.FAIL(code, nil))
		}
	}
}
func _queryRepository(c *gin.Context) {
	if page, err := util.GetParamJson[util.Page[*Repository]](c); err == nil {
		if err := queryList(&page); err == nil {
			c.JSON(http.StatusOK, util.SUCCESS(page))
			return
		}
		c.JSON(http.StatusOK, util.FAIL(util.MsgCodeFail, err))
	}
}
func getRepositoryByName(name string) (Repository, error) {
	if repos, ok := Store[name]; ok {
		return *repos, nil
	} else {
		return *repos, fmt.Errorf(util.MsgCodeReposNotExists)
	}
}
func _viewComponent(c *gin.Context) {
	reposName := c.Param("libName")
	if repos, err := getRepositoryByName(reposName); err == nil {
		if compo, err := util.GetParamJson[Component](c); err == nil {
			c.JSON(http.StatusOK, repos.GetComponentList(compo.Path))
		} else {
			c.JSON(http.StatusOK, util.FAIL(err.Error(), nil))
		}
	} else {
		c.JSON(http.StatusOK, util.FAIL(err.Error(), nil))
	}
}
func _deleteComponent(c *gin.Context) {
	var (
		err   error
		repos Repository
		compo Component
	)
	reposName := c.Param("libName")
	if repos, err = getRepositoryByName(reposName); err == nil {
		if compo, err = util.GetParamJson[Component](c); err == nil {
			if err = repos.delComponent(&compo); err == nil {
				c.JSON(http.StatusOK, util.SUCCESS(nil))
				return
			}
		}
	}
	c.JSON(http.StatusOK, util.FAIL(err.Error(), nil))

}
func _uploadComponent(c *gin.Context) {
	var (
		err        error
		repository *Repository
		length     int
		data       []byte
	)
	length, err = strconv.Atoi(c.GetHeader("Content-Length"))
	if length <= 0 {
		err = fmt.Errorf("invalid content length")
	}
	if err == nil {
		if data, err = io.ReadAll(c.Request.Body); err == nil {
			if length != len(data) {
				err = fmt.Errorf("invalid content length not equal to data length")
			}
			if err == nil {
				if repository, err = checkAndGetRepository(c); err == nil {
					if repository.Mode&2 != 2 {
						err = fmt.Errorf("invalid repository mode is not support write")
					}
					if err == nil {
						filePath := c.Param("filePath")
						localFilePath := repository.GetComponent(filePath)
						if err = saveFile(localFilePath, data); err == nil {
							if err = util.GenerateHash(localFilePath); err != nil {
								err = fmt.Errorf("generate hash error: %v", err)
							} else {
								c.JSON(http.StatusOK, util.SUCCESS(nil))
								return
							}
						}
					}
				}
			}
		}
	}
	c.JSON(http.StatusOK, util.FAIL(err.Error(), nil))
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

	localFilePath := repository.GetComponent(filePath)

	f, err := fs.Open(localFilePath)
	defer closeFile(f)

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
		if !checkAuthPublish(c) {
			c.String(http.StatusUnauthorized, "Unauthorised")
			return
		}

		if err = util.GenerateHash(localFilePath); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("generate hash failed, message: %v\n", err))
		}
	}

	u := repository.GetComponent(filePath)
	c.Request.URL.RawPath = u
	c.Request.URL.Path = u

	fileServer.ServeHTTP(c.Writer, c.Request)
}

func put(c *gin.Context) {
	if !checkAuthPublish(c) {
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
	localFilePath := repository.GetComponent(filePath)
	if err = saveFile(localFilePath, data); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("write file failed. message: %v\n", err))
		return
	}

	if generate := c.Query("generate_md5_sha1"); strings.EqualFold(generate, "true") {
		if err = util.GenerateHash(localFilePath); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("generate hash failed, message: %v\n", err))
		}
	}

	c.String(http.StatusOK, "OK")
}
