package util

import (
	"crypto/rsa"
	"embed"
	"flag"
	"fmt"
	"github.com/creasty/defaults"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	Log        = logrus.New()
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
	Engine            *gin.Engine
	fs                http.FileSystem
	fileServer        http.Handler
	client            = resty.New()
	Static            embed.FS
	AuthHandler       func(c *gin.Context) bool
)

func init() {
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	initConfig()
	initLog()
	initKey()
	initAuth()
	initRouter()
}
func initRouter() {
	//处理路由信息
	fs = http.Dir(config.LocalRepository)
	fileServer = http.StripPrefix(path.Join("/", config.Context), http.FileServer(fs))
	gin.SetMode(gin.ReleaseMode)
	Engine = gin.Default()
	Engine.Use(GinLogger())
	Engine.Use(jwtToken())
}
func initAuth() {
	authExclude := config.AuthExclude
	authExclude = strings.Replace(authExclude, "*", ".*", -1)
	authExclude = strings.Replace(authExclude, ",", ")|(", -1)
	authExclude = fmt.Sprintf("^((%s))(\\?.*)?$", authExclude)
	authExcludeRegexp = regexp.MustCompile(authExclude)
}
func initKey() {
	KeyDir = filepath.Dir(config.DataDir)
	// rsa公钥私钥处理
	privateKey, PublicKey = rsaGenerate(1024)
}
func initConfig() {
	var err error

	// 命令行参数解析
	flag.StringVar(&configPath, "c", "config.yaml", "配置文件路径")
	flag.Parse()

	Log.Infof("configure file: %s", configPath)
	// 读取配置文件
	if file, err = os.ReadFile(configPath); err != nil {
		Log.Errorf("config.yaml read error %v", err)
	}
	// 解析yaml
	if err = yaml.Unmarshal(file, config); err != nil {
		Log.Errorf("config.yaml unmarshal error %v", err)
	}
	// 添加默认值
	if err = defaults.Set(config); err != nil {
		Log.Errorf("set defaults error %v", err)
	}
}

func initLog() {
	Log.SetLevel(config.Logging.Level)
	if config.Logging.Path != "" {
		logFile := path.Join(config.Logging.Path, config.Logging.Level.String()+".log")
		if err := CreateFileIfNotExist(logFile); err != nil {
			Log.Errorf("create log file error, file is: %s, message: %v", logFile, err)
			return
		}
		src, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			Log.Errorf("open log file error, file is: %s, message: %v", logFile, err)
			return
		}
		// 同时写文件和屏幕
		fileAndStdoutWriter := io.MultiWriter(src, os.Stdout)
		Log.SetOutput(fileAndStdoutWriter)
	}
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		Log.Infof("| %3d | %13v | %15s | %s | %s",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
