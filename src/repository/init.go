package repository

import (
	"Mini-Repository/src/util"
	"github.com/go-resty/resty/v2"
	"path"
)

var (
	list       []*Repository
	config     = util.LoadConfig()
	fs         = util.GetFileSystem()
	log        = util.Log
	Store      map[string]*Repository
	cache      map[int]*Repository
	client     = resty.New()
	fileServer = util.GetFileServer()
	dataPath   = path.Join(config.DataDir, "repository.json")
	maxId      = 0
)

func init() {
	list = loadFile()
	Store = make(map[string]*Repository, len(list))
	cache = make(map[int]*Repository, len(list))
	for _, v := range list {
		Store[v.Name] = v
		cache[v.Id] = v
		if maxId < v.Id {
			maxId = v.Id
		}
	}
}
