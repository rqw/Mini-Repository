package user

import (
	"Mini-Repository/src/permission"
	"Mini-Repository/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterRegister() {
	util.Engine.GET("/user/:id", _getUserInfoById)
	util.Engine.DELETE("/user/:id", _dropUserById)
	util.Engine.PUT("/user", _saveUserInfo)
	util.Engine.POST("/user/passwd", _passwd)
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
			permList := permission.GetPermissionList(user.PermissionList)
			roleList := make([]Role, len(permList))
			for i, perm := range permList {
				roleList[i] = Role{
					Value:    perm.Name,
					RoleName: perm.Name,
				}
			}
			user.Roles = roleList
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

func _saveUserInfo(c *gin.Context) {
	if user, err := util.GetParamJson[User](c); err == nil {
		code := saveUser(&user)
		if code == util.MsgCodeSuccess {
			c.JSON(http.StatusOK, util.SUCCESS(user))
		} else {
			c.JSON(http.StatusOK, util.FAIL(code, nil))
		}
	}
}

func _queryUser(c *gin.Context) {
	if page, err := util.GetParamJson[util.Page[*User]](c); err == nil {
		if err := queryList(&page); err == nil {
			c.JSON(http.StatusOK, util.SUCCESS(page))
			return
		}
		c.JSON(http.StatusOK, util.FAIL(util.MsgCodeFail, err))
	}
}
func _passwd(c *gin.Context) {
	if passwdInfo, err := util.GetParamJson[PasswdInfo](c); err == nil {
		o, _ := c.Get("current")
		user := o.(User)
		passwdInfo.ID = user.ID
		if modifyPassword(&passwdInfo) {
			c.JSON(http.StatusOK, util.SUCCESS(nil))
		} else {
			c.JSON(http.StatusOK, util.FAIL(util.MsgCodeFail, nil))
		}
	}
}
