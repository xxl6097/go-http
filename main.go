package main

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server"
	"github.com/xxl6097/go-http/server/config"
)

type ConfigEx struct {
	config.Config
	Version string `yaml:"version"`
}

func init() {
	//config.Init(&ConfigEx{})
	glog.SetLogFile("./log", "app.log")
	glog.SetCons(true)
}

func bootstrap() {
	glog.Errorf("config--->%+v", config.Get().GetConfig())
	server.NewServer().Start(fmt.Sprintf(":%d", config.Get().GetConfig().HttpConfig.Server.Port))
}

func main() {
	bootstrap()
}
