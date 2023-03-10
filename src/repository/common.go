package repository

import (
	"Mini-Repository/src/user"
	"Mini-Repository/src/util"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"os"
	"path"
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

func checkAuthPublish(c *gin.Context) bool {
	libName := c.Param("libName")

	authorization := c.GetHeader("Authorization")
	if !strings.HasPrefix(authorization, "Basic ") {
		return false
	}
	// 校验用户
	authorization = strings.TrimSpace(authorization[6:])
	if data, err := base64.StdEncoding.DecodeString(authorization); err == nil {
		str := string(data)
		infos := strings.Split(str, ":")
		if len(infos) == 2 {
			return user.AuthRepos(infos[0], infos[1], libName)
		}

	}

	return false
}

func (repos Repository) GetComponent(filePath string) string {
	if repos.DiskPath != "" {
		return path.Join(repos.DiskPath, filePath)
	}
	return path.Join(config.Context, repos.Name, filePath)
}
func (repos Repository) GetComponentList(filePath string) []*Component {
	list, err := os.ReadDir(repos.GetComponent(filePath))
	if err != nil && len(list) == 0 {
		return []*Component{}
	}
	components := make([]*Component, len(list))
	for _, f := range list {
		component := Component{
			Name:  f.Name(),
			IsDir: f.IsDir(),
		}
		if f.IsDir() {
			if info, err := f.Info(); err != nil {
				component.Size = info.Size()
				component.ModTime = info.ModTime().UnixMilli()
			}
		}
		components = append(components, &component)
	}
	return components
}
func (repos Repository) delComponent(component *Component) error {
	if component.IsDir {
		return os.RemoveAll(component.Path)
	} else {
		return os.Remove(component.Path)
	}
}
