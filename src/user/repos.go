package user

import (
	"Mini-Repository/src/util"
	"encoding/json"
	"os"
	"path"
)

var (
	UserCache      map[int]*User
	LoginNameCache map[string]*User
	UserList       []*User
	MaxId          = 0
	log            = util.Log
	config         = util.LoadConfig()
	dataPath       = path.Join(config.DataDir, "user.json")
)

func init() {
	UserList = loadFile()
	UserCache = make(map[int]*User, len(UserList))
	LoginNameCache = make(map[string]*User, len(UserList))
	for _, user := range UserList {
		UserCache[user.ID] = user
		LoginNameCache[user.LoginName] = user
		if user.ID > MaxId {
			MaxId = user.ID
		}
	}
}
func findUserInfoById(id int) User {
	return *UserCache[id]
}
func findAllUser() []*User {
	return UserList
}
func validaUser(user User) bool {
	if localUser, state := LoginNameCache[user.LoginName]; state {
		return localUser.Password == user.Password
	}
	return false
}
func delUserById(id int) {
	user := UserCache[id]
	delete(LoginNameCache, user.LoginName)
	delete(UserCache, id)
	var index int
	for i, u := range UserList {
		if user.ID == u.ID {
			index = i
			break
		}
	}
	UserList = append(UserList[:index], UserList[index+1:]...)
	saveToFile(UserList)
}
func addUser(user *User) string {
	if _, state := LoginNameCache[user.LoginName]; !state {
		return util.MsgCodeUserExists
	}
	if user.Password == "" {
		return util.MsgCodeUserPwdNotEmpty
	}
	user.ID = MaxId + 1
	MaxId++
	UserCache[user.ID] = user
	LoginNameCache[user.LoginName] = user
	UserList[len(UserList)] = user
	saveToFile(UserList)
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
	admin := User{ID: 1, LoginName: "admin", Password: string(util.GetHash([]byte("admin"), "md5")), Fullname: "admin", PermissionList: []int{0}}
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
