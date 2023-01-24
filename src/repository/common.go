package repository

import (
	"Mini-Repository/src/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"os"
	"strings"
)

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
	repository = Store[libName]
	if repository == nil || repository.Mode == 0 {
		return nil, errors.New(fmt.Sprintf("repository %s is not actived", libName))
	}

	return repository, nil
}
func closeFile(f http.File) {
	if f != nil {
		_ = f.Close()
	}
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

func saveFile(localFilePath string, data []byte) error {
	if err := util.CreateParentIfNotExist(localFilePath); err != nil {
		return err
	}

	if err := os.WriteFile(localFilePath, data, 0755); err != nil {
		return err
	}
	return nil
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
