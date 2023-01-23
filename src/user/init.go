package user

import "Mini-Repository/src/util"

func init() {
	list = loadFile()
	cache = make(map[int]*User, len(list))
	LoginNameCache = make(map[string]*User, len(list))
	for _, user := range list {
		cache[user.ID] = user
		LoginNameCache[user.LoginName] = user
		if user.ID > maxId {
			maxId = user.ID
		}
	}
	util.FindUserInfoById = _FindUserInfoById
}
