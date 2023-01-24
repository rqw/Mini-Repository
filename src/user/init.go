package user

import (
	"Mini-Repository/src/permission"
	"Mini-Repository/src/util"
	"encoding/base64"
	"fmt"
	"path"
	"regexp"
)

var (
	cache          map[int]*User
	LoginNameCache map[string]*User
	userAuthCache  map[int]*regexp.Regexp

	list     []*User
	maxId    = 0
	log      = util.Log
	config   = util.LoadConfig()
	dataPath = path.Join(config.DataDir, "user.json")
)

func init() {
	list = loadFile()
	cache = make(map[int]*User, len(list))
	LoginNameCache = make(map[string]*User, len(list))
	userAuthCache = make(map[int]*regexp.Regexp, len(list))
	for _, user := range list {
		cache[user.ID] = user
		LoginNameCache[user.LoginName] = user
		userAuthCache[user.ID] = permission.CompileToRegexp(user.PermissionList)
		if user.ID > maxId {
			maxId = user.ID
		}
	}
	util.AuthHandler = _AuthHandler
	for _, user := range list {
		base := fmt.Sprintf("%s:%s", user.LoginName, user.Mrt)
		auth := base64.StdEncoding.EncodeToString([]byte(base))
		config.Auth[auth] = auth
	}
}
