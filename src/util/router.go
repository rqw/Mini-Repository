package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterRegister() {
	Engine.Any("/ui/*filepath", staticFs)
	// 匹配vue中的/v/*链接，跳转至vue入口文件，vue会自动进行路由
	Engine.GET("/ui", getUi)
	// 匹配/链接，重定向到主页
	Engine.GET("/", firstPage)
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

func jwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		curl := c.Request.RequestURI
		if authExcludeRegexp.MatchString(curl) {
			c.Next()
		} else {
			if AuthHandler(c) {
				c.Next()
			} else {
				c.Abort()
			}
		}
	}
}
