package main

import (
	"Mini-Repository/src/permission"
	"Mini-Repository/src/repository"
	"Mini-Repository/src/user"
	"Mini-Repository/src/util"
	"embed"
	"encoding/json"
	"fmt"
)

var log = util.Log

//go:embed ui/*
var Static embed.FS

func main() {
	config := util.LoadConfig()
	util.Static = Static
	util.RouterRegister()
	user.RouterRegister()
	repository.RouterRegister()
	permission.RouterRegister()
	if bytes, err := json.Marshal(config); err == nil {
		log.Debugf("启动参数: %s", bytes)
	} else {
		log.Panic(err)
	}

	if err := util.Engine.Run(fmt.Sprintf("%s:%s", config.Listen, config.Port)); err != nil {
		log.Errorln("服务启动失败")
	}
}
