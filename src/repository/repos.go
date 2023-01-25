package repository

import (
	"Mini-Repository/src/permission"
	"Mini-Repository/src/util"
	"fmt"
)

func SaveRepository(repos *Repository) string {
	if _, state := Store[repos.Name]; !state {
		return util.MsgCodeReposExists
	}

	if o, state := cache[repos.Id]; !state {
		repos.Id = maxId + 1
		maxId++
		Store[repos.Name] = repos
		cache[repos.Id] = repos
		list = append(list, repos)
		readPerm := permission.Permission{Type: permission.TypeRepos, Name: fmt.Sprintf("%s-read", repos.Name), Description: fmt.Sprintf("%s Repository read permission.", repos.Name), ValidRegex: fmt.Sprintf("POST@/repository/view/%s", repos.Name)}
		writePerm := permission.Permission{Type: permission.TypeRepos, Name: fmt.Sprintf("%s-write", repos.Name), Description: fmt.Sprintf("%s Repository write permission.", repos.Name), ValidRegex: fmt.Sprintf("POST@/repository/upload/%s", repos.Name)}
		deletePerm := permission.Permission{Type: permission.TypeRepos, Name: fmt.Sprintf("%s-delete", repos.Name), Description: fmt.Sprintf("%s Repository delete permission.", repos.Name), ValidRegex: fmt.Sprintf("POST@/repository/del/%s", repos.Name)}
		publishPerm := permission.Permission{Type: permission.TypeRepos, Name: fmt.Sprintf("%s-publish", repos.Name), Description: fmt.Sprintf("%s Repository publish permission.", repos.Name), ValidRegex: fmt.Sprintf("REPOS@PUBLISH@%s", repos.Name)}
		permission.AddPermissionList([]*permission.Permission{&readPerm, &writePerm, &deletePerm, &publishPerm})
	} else {
		o.Mode = repos.Mode
		o.PublicAccess = repos.PublicAccess
		o.Cache = repos.Cache
		o.Mirror = repos.Mirror
	}
	saveToFile(list)
	return util.MsgCodeSuccess
}
func DelRepository(id int) {
	var index int
	for i, u := range list {
		if id == u.Id {
			index = i
			delete(Store, u.Name)
			delete(cache, u.Id)
			break
		}
	}
	repos := list[index]
	list = append(list[:index], list[index+1:]...)
	permission.DelPermissionList([]string{fmt.Sprintf("%s-read", repos.Name), fmt.Sprintf("%s-write", repos.Name), fmt.Sprintf("%s-delete", repos.Name), fmt.Sprintf("%s-publish", repos.Name)})
	saveToFile(list)
}

func saveToFile(list []*Repository) {
	if err := util.AnyToJsonFile(list, dataPath); err != nil {
		log.Errorf("saveToFile repository.json fail.%v", err)
	}
}
func loadFile() []*Repository {
	var cache []*Repository
	if err := util.JsonFileToAny(dataPath, &cache); err != nil {
		log.Errorf("loadFile repository.json fail.%v", err)
		cache = make([]*Repository, 0)
		saveToFile(cache)
	}
	return cache
}
