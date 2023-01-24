package repository

import (
	"Mini-Repository/src/util"
)

func GetAllRepository() []*Repository {
	return list
}
func AddRepository(repos *Repository) string {
	if _, state := Store[repos.Name]; !state {
		return util.MsgCodeReposExists
	}
	repos.Id = maxId + 1
	maxId++
	Store[repos.Name] = repos
	cache[repos.Id] = repos
	list = append(list, repos)
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
	list = append(list[:index], list[index+1:]...)
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
