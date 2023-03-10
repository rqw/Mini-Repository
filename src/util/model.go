package util

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Listen          string        `yaml:"listen" default:"0.0.0.0"`
	Port            string        `yaml:"port" default:"8888"`
	Context         string        `yaml:"context" default:"repos"`
	LocalRepository string        `yaml:"localRepository" default:"./repos"`
	AuthExclude     string        `yaml:"authExclude" default:"/,/ui,/repos/?*,/user/login"`
	DataDir         string        `yaml:"dataDir" default:"./data"`
	User            []*User       `yaml:"user" default:"[{\"Name\":\"user\",\"Password\":\"password\"}]"`
	Repository      []*Repository `yaml:"repository" default:"[{\"Id\":\"public\",\"Name\":\"mirror\",\"Mirror\":[\"https://repo1.maven.org/maven2\",\"https://maven.aliyun.com/nexus/content/repositories/public\"]}]"`
	Logging         *Logging      `yaml:"logging" default:"{\"Path\":\"./logs\",\"Level\":\"info\"}"`
	Auth            map[string]interface{}
	RepositoryStore map[string]*Repository
}

type User struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

type Repository struct {
	Id     string   `yaml:"id"`
	Name   string   `yaml:"name"`
	Target string   `yaml:"target"`
	Mode   int      `yaml:"mode" default:"4"`
	Cache  bool     `yaml:"cache" default:"false"`
	Mirror []string `yaml:"mirror"`
}

type Logging struct {
	Path  string       `yaml:"path" default:""`
	Level logrus.Level `yaml:"level" default:"debug"`
}
type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
type Page[T any] struct {
	No        int               `json:"no"`        //当前页码，从1开始
	Total     int               `json:"total"`     //记录总数
	Condition map[string]string `json:"condition"` //查询条件
	Orders    []string          `json:"orders"`    //排序
	Capacity  int               `json:"capacity"`  //页容量
	DataList  []T               `json:"dataList"`  //数据列表
}

func (page Page[any]) GetFirst() int {
	return (page.No - 1) * page.Capacity
}

func SUCCESS(data interface{}) Resp {
	return Resp{Code: 0, Result: data}
}
func FAIL(message string, data interface{}) Resp {
	return Resp{Code: -1, Message: message, Result: data}
}
