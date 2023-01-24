package user

import (
	"Mini-Repository/src/permission"
	"Mini-Repository/src/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
func validaUser(user *User) bool {
	if localUser, state := LoginNameCache[user.LoginName]; state {
		if localUser.Password == user.Password {
			t, _ := json.Marshal(localUser)
			json.Unmarshal(t, &user)
			user.Password = ""
			user.JwtToken = util.ReleaseToken(user.ID, user.Act, 12*60*60)
			return true
		}
	}
	return false
}
func delUserById(id int) {
	user := cache[id]
	delete(LoginNameCache, user.LoginName)
	delete(cache, id)
	delete(userAuthCache, id)
	var index int
	for i, u := range list {
		if user.ID == u.ID {
			index = i
			break
		}
	}
	list = append(list[:index], list[index+1:]...)
	base := fmt.Sprintf("%s:%s", user.LoginName, user.Mrt)
	auth := base64.StdEncoding.EncodeToString([]byte(base))
	delete(config.Auth, auth)
	saveToFile(list)
}
func addUser(user *User) string {
	if _, state := LoginNameCache[user.LoginName]; !state {
		return util.MsgCodeUserExists
	}
	if user.Password == "" {
		return util.MsgCodeUserPwdNotEmpty
	}
	user.ID = maxId + 1
	user.Act = util.Md5(strconv.FormatInt(time.Now().UnixNano(), 10))
	user.Mrt = util.Md5(fmt.Sprintf("%d-%s-%s-%d", user.ID, user.LoginName, user.Password, time.Now().UnixNano()))
	maxId++
	userAuthCache[user.ID] = permission.CompileToRegexp(user.PermissionList)
	cache[user.ID] = user
	LoginNameCache[user.LoginName] = user
	list = append(list, user)
	base := fmt.Sprintf("%s:%s", user.LoginName, user.Mrt)
	auth := base64.StdEncoding.EncodeToString([]byte(base))
	config.Auth[auth] = auth
	saveToFile(list)
	return util.MsgCodeSuccess
}

func defaultPermission() []*User {
	admin := User{ID: 1, LoginName: "admin", Password: util.Md5("admin"), Fullname: "admin", PermissionList: []int{1, 2, 3}, Act: util.Md5(strconv.FormatInt(time.Now().UnixNano(), 10))}
	admin.Mrt = util.Md5(fmt.Sprintf("%d-%s-%s-%d", admin.ID, admin.LoginName, admin.Password, time.Now().UnixNano()))
	return []*User{&admin}
}
func saveToFile(list []*User) {
	if err := util.AnyToJsonFile(list, config.DataDir); err != nil {
		log.Errorf("saveToFile user.json fail.%v", err)
	}
}
func loadFile() []*User {
	var list []*User
	if err := util.JsonFileToAny(dataPath, &list); err != nil {
		log.Errorf("loadFile user.json fail.%v", err)
		list = defaultPermission()
		saveToFile(list)
	}
	return list
}
