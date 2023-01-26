package permission

import (
	"Mini-Repository/src/util"
	"fmt"
	"regexp"
	"strings"
)

func GetAllPermission() []*Permission {
	return list
}
func AddPermission(perm *Permission) {
	perm.ID = maxId + 1
	maxId++
	cache[perm.ID] = perm
	list = append(list, perm)
	saveToFile(PermissionCache{MaxId: maxId, PermissionList: list})
}
func AddPermissionList(permList []*Permission) {
	for _, perm := range permList {
		perm.ID = maxId + 1
		maxId++
		cache[perm.ID] = perm
		list = append(list, perm)
	}
	saveToFile(PermissionCache{MaxId: maxId, PermissionList: list})
}

func DelPermission(name string) {
	var index int
	for i, u := range list {
		if name == u.Name {
			index = i
			delete(cache, u.ID)
			break
		}
	}
	list = append(list[:index], list[index+1:]...)
	saveToFile(PermissionCache{MaxId: maxId, PermissionList: list})
}

func DelPermissionList(nameList []string) {
	for _, name := range nameList {
		var index int
		for i, u := range list {
			if name == u.Name {
				index = i
				delete(cache, u.ID)
				break
			}
		}
		list = append(list[:index], list[index+1:]...)
	}
	saveToFile(PermissionCache{MaxId: maxId, PermissionList: list})
}

func GetPermissionList(ids []int) []Permission {
	permList := make([]Permission, len(ids))
	for i, id := range ids {
		if perm, s := cache[id]; s {
			permList[i] = *perm
		}
	}
	return permList
}
func CompileToRegexp(ids []int) *regexp.Regexp {
	validRegexList := make([]string, len(ids))
	for i, id := range ids {
		if perm, s := cache[id]; s {
			validRegexList[i] = perm.ValidRegex
		}
	}
	regexStr := fmt.Sprintf("^((%s))(\\?.*)?$", strings.Join(validRegexList, ")|("))
	return regexp.MustCompile(regexStr)
}

func defaultUser() PermissionCache {
	p1 := Permission{ID: 1, Name: "USER-MANAGER", Type: TypeMenu, Description: "管理系统用户维护的权限，谨慎授予", ValidRegex: "(.*@/user(/[^/]+)?)|(POST@/permission)"}
	p2 := Permission{ID: 2, Name: "REPOSITORY-MANAGER", Type: TypeMenu, Description: "仓库信息维护的权限，谨慎授予", ValidRegex: ".*@/repository(/[^/]+)?"}
	p3 := Permission{ID: 3, Name: "REPOSITORY-VIEW", Type: TypeMenu, Description: "仓库浏览权限，可以查看仓库内容", ValidRegex: "POST@/repository(/[^/]+)?"}

	return PermissionCache{MaxId: 3, PermissionList: []*Permission{&p1, &p2, &p3}}
}
func saveToFile(list PermissionCache) {
	if err := util.AnyToJsonFile(list, dataPath); err != nil {
		log.Errorf("saveToFile permission.json fail.%v", err)
	}
}
func loadFile() PermissionCache {
	var cache PermissionCache
	if err := util.JsonFileToAny(dataPath, &cache); err != nil {
		log.Errorf("loadFile permission.json fail.%v", err)
		cache = defaultUser()
		saveToFile(cache)
	}
	return cache
}
