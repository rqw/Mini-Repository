package permission

import (
	"Mini-Repository/src/util"
	"path"
)

var (
	cache    map[int]*Permission
	list     []*Permission
	maxId    = 0
	log      = util.Log
	config   = util.LoadConfig()
	dataPath = path.Join(config.DataDir, "permission.json")
)

func init() {
	t := loadFile()
	maxId = t.MaxId
	cache = make(map[int]*Permission, len(t.PermissionList))
	list = t.PermissionList
	for _, perm := range list {
		cache[perm.ID] = perm
	}
}
