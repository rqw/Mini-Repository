package user

import (
	"Mini-Repository/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterRegister() {
	util.Engine.GET("/user/:id", _getUserInfoById)
	util.Engine.DELETE("/user/:id", _dropUserById)
	util.Engine.PUT("/user", _addUserInfo)
	util.Engine.POST("/user", _queryUser)
	util.Engine.POST("/user/login", _login)
}
func _login(c *gin.Context) {
	if user, err := util.GetParamJson[User](c); err == nil {
		if validaUser(&user) {
			c.JSON(http.StatusOK, util.SUCCESS(user))
			return
		}
		c.JSON(http.StatusOK, util.FAIL(util.MsgCodeUserValidFail, nil))
	}
}
func _getUserInfoById(c *gin.Context) {
	if id, err := util.GetParamId(c); err == nil {
		if id == 0 {
			o, _ := c.Get("current")
			user := o.(User)
			user.Password = ""
			c.JSON(http.StatusOK, util.SUCCESS(user))
			return
		}
		user := FindUserInfoById(id)
		c.JSON(http.StatusOK, util.SUCCESS(user))
	}
}
func _dropUserById(c *gin.Context) {
	if id, err := util.GetParamId(c); err == nil {
		delUserById(id)
		c.JSON(http.StatusOK, util.SUCCESS(nil))
	}
}

func _addUserInfo(c *gin.Context) {
	if user, err := util.GetParamJson[User](c); err == nil {
		code := addUser(&user)
		if code == util.MsgCodeSuccess {
			c.JSON(http.StatusOK, util.SUCCESS(nil))
		} else {
			c.JSON(http.StatusOK, util.FAIL(code, nil))
		}
	}
}

func _queryUser(c *gin.Context) {
	c.JSON(http.StatusOK, util.SUCCESS(list))
}
