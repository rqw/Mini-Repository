package user

import (
	"Mini-Repository/src/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterRegister() {
	util.Engine.GET("/user/:id", _getUserInfoById)
	util.Engine.DELETE("/user/:id", _dropUserById)
	util.Engine.PUT("/user", _addUserInfo)
	util.Engine.POST("/user", _queryUser)
	util.Engine.POST("/user/login", _login)
}
func _login(c *gin.Context) {
	if user, err := _getParamUser(c); err == nil {
		if validaUser(user) {
			c.JSON(http.StatusOK, util.SUCCESS(user))
			return
		}
		c.JSON(http.StatusOK, util.FAIL(util.MsgCodeUserValidFail, nil))
	}
}
func _getUserInfoById(c *gin.Context) {
	if id, err := _getParamId(c); err == nil {
		user := findUserInfoById(id)
		c.JSON(http.StatusOK, util.SUCCESS(user))
	}
}
func _dropUserById(c *gin.Context) {
	if id, err := _getParamId(c); err == nil {
		delUserById(id)
		c.JSON(http.StatusOK, util.SUCCESS(nil))
	}
}

func _addUserInfo(c *gin.Context) {
	if user, err := _getParamUser(c); err == nil {
		code := addUser(&user)
		if code == util.MsgCodeSuccess {
			c.JSON(http.StatusOK, util.SUCCESS(nil))
		} else {
			c.JSON(http.StatusOK, util.FAIL(code, nil))
		}
	}
}

func _queryUser(c *gin.Context) {
	c.JSON(http.StatusOK, util.SUCCESS(findAllUser()))
}
func _getParamUser(c *gin.Context) (User, error) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, util.FAIL(err.Error(), nil))
	}
	return user, err
}
func _getParamId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, util.FAIL(err.Error(), nil))
	}
	return id, err
}
