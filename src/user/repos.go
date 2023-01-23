package user

import (
	"Mini-Repository/src/util"
	"encoding/json"
	"os"
	"path"
	"strconv"
	"time"
)

var (
	cache          map[int]*User
	LoginNameCache map[string]*User
	list           []*User
	maxId          = 0
	log            = util.Log
	config         = util.LoadConfig()
	dataPath       = path.Join(config.DataDir, "user.json")
)

func FindUserInfoById(id int) User {
	user, _ := cache[id]
	return *user
}

func _FindUserInfoById(id int) any {
	return FindUserInfoById(id)
}
func findAllUser() []*User {
	return list
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
	var index int
	for i, u := range list {
		if user.ID == u.ID {
			index = i
			break
		}
	}
	list = append(list[:index], list[index+1:]...)
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
	maxId++
	cache[user.ID] = user
	LoginNameCache[user.LoginName] = user
	list[len(list)] = user
	saveToFile(list)
	return util.MsgCodeSuccess
}
func saveToFile(list []*User) {
	content, err := json.Marshal(list)
	if err != nil {
		log.Errorf("saveToFile user.json json.Marshal fail.%v", err)
		return
	}
	if isExists, _ := util.CheckFileExist(config.DataDir); !isExists {
		os.MkdirAll(config.DataDir, os.ModePerm)
	}
	err = os.WriteFile(dataPath, content, os.ModePerm)
	if err != nil {
		log.Errorf("saveToFile user.json WriteFile fail.%v", err)
	}

}
func defaultUser() []*User {
	admin := User{ID: 1, LoginName: "admin", Password: util.Md5("admin"), Fullname: "admin", PermissionList: []int{0}, Act: util.Md5(strconv.FormatInt(time.Now().UnixNano(), 10))}
	return []*User{&admin}
}
func loadFile() []*User {
	var list []*User
	if isExists, _ := util.CheckFileExist(dataPath); !isExists {
		data, err := os.ReadFile(dataPath)
		if err == nil {
			err = json.Unmarshal(data, &list)
			if err == nil {
				return list
			}
		}
		if err != nil {
			log.Errorf("loadFile user.json fail.%v", err)
		}
	}
	if list == nil {
		list = defaultUser()
		saveToFile(list)
	}
	return list
}
