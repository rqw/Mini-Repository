package util

import (
	"crypto/rsa"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	log        = logrus.New()
	file       []byte
	configPath string
	KeyDir     string
	PublicKey  rsa.PublicKey
	privateKey *rsa.PrivateKey
	config     = &Config{
		Auth:            make(map[string]interface{}),
		RepositoryStore: make(map[string]*Repository),
	}
	authExcludeRegexp *regexp.Regexp
)

func init() {
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	// 设置公共目录信息
	KeyDir = filepath.Dir(config.DataDir)
	var err error

	// 命令行参数解析
	flag.StringVar(&configPath, "c", "config.yaml", "配置文件路径")
	flag.Parse()

	log.Infof("configure file: %s", configPath)
	// 读取配置文件
	if file, err = os.ReadFile(configPath); err != nil {
		log.Errorf("config.yaml read error %v", err)
	}
	// 解析yaml
	if err = yaml.Unmarshal(file, config); err != nil {
		log.Errorf("config.yaml unmarshal error %v", err)
	}
	// 添加默认值
	if err = defaults.Set(config); err != nil {
		log.Errorf("set defaults error %v", err)
	}
	authExclude := config.AuthExclude
	authExclude = strings.Replace(authExclude, "*", ".*", -1)
	authExclude = strings.Replace(authExclude, ",", ")|(", -1)
	authExclude = fmt.Sprintf("^((%s))(\\?.*)?$", authExclude)
	log.Infof("auth exclude: %s", authExclude)
	authExcludeRegexp = regexp.MustCompile(authExclude)
	log.Infof("auth exclude regexp: %s", authExcludeRegexp)
	// rsa公钥私钥处理
	privateKey, PublicKey = rsaGenerate(1024)
	// 预处理认证信息
	for _, user := range config.User {
		base := fmt.Sprintf("%s:%s", user.Name, user.Password)
		auth := base64.StdEncoding.EncodeToString([]byte(base))
		config.Auth[auth] = auth
	}
	// 预处理存储库
	for _, repository := range config.Repository {
		// 移除未启用的repository
		if repository.Mode == 0 {
			continue
		}
		// 如果没设置目标目录, 则默认使用Id
		if repository.Target == "" {
			repository.Target = repository.Id
		}
		config.RepositoryStore[repository.Id] = repository
		log.Infof("repository: http://%s:%s/%s/repos/%s local dirname: %s", config.Listen, config.Port, config.Context, repository.Id, repository.Target)
	}
	//处理路由信息
	fs = http.Dir(config.LocalRepository)
	fileServer = http.StripPrefix(path.Join("/", config.Context), http.FileServer(fs))
	gin.SetMode(gin.ReleaseMode)
	Engine = gin.Default()
	Engine.Use(GinLogger())
	Engine.Use(jwtToken())
}
