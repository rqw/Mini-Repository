package user

import (
	"Mini-Repository/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindUserInfoById(id int) User {
	user, _ := cache[id]
	return *user
}

func _AuthHandler(c *gin.Context) bool {
	url := fmt.Sprintf("%s@%s", c.Request.Method, c.Request.RequestURI)
	token := c.Request.Header.Get("Authorization")
	if token != "" {
		state, id, act := util.ValidToken(token)
		if state {
			//有权限继续执行,将当前用户信息存入上下文
			if re, s := userAuthCache[id]; s && re.MatchString(url) {
				currentUser := FindUserInfoById(id)
				c.Set("id", id)
				c.Set("act", act)
				c.Set("current", currentUser)
				return true
			} else {
				c.JSON(http.StatusOK, util.FAIL(util.MsgCodeAuthFail, nil))
			}
		}
	} else {
		c.JSON(http.StatusOK, util.FAIL(util.MsgCodeTokenValidFail, nil))
	}
	return false
}
func AuthRepos(loginName string, mrt string, libName string) bool {
	if user, ok := LoginNameCache[loginName]; ok && user.Mrt == mrt {
		if re, s := userAuthCache[user.ID]; s && re.MatchString(fmt.Sprintf("REPOS@PUBLISH@%s", libName)) {
			return true
		}
	}
	return false
}
