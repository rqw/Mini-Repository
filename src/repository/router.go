package repository

import (
	"Mini-Repository/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func RouterRegister() {
	util.Engine.GET("/repository/:id", _getRepository)
	util.Engine.DELETE("/repository/:id", _dropRepository)
	util.Engine.PUT("/repository", _saveRepository)
	util.Engine.POST("/repository", _queryRepository)
	util.Engine.POST("/repository/view/:libName", _viewRepository)
	util.Engine.POST("/repository/del/:libName", _deleteRepository)
	util.Engine.POST("/repository/upload/:libName", _uploadRepository)
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
	c.JSON(http.StatusOK, util.SUCCESS(list))
}
func _viewRepository(c *gin.Context) {
	// todo: implement
}
func _deleteRepository(c *gin.Context) {
	// todo: implement
}
func _uploadRepository(c *gin.Context) {
	// todo: implement
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

	localFilePath := path.Join(repository.DiskPath, filePath)

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
		if !checkAuthPublish(c) {
			c.String(http.StatusUnauthorized, "Unauthorised")
			return
		}

		if err = util.GenerateHash(localFilePath); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("generate hash failed, message: %v\n", err))
		}
	}

	u := fmt.Sprintf("/%s/%s%s", config.Context, repository.DiskPath, filePath)
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
	localFilePath := path.Join(config.LocalRepository, repository.DiskPath, filePath)
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
