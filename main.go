package main

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"go-http/server"
	"go-http/server/config"
)

var conf config.Yml

func init() {
	conf = config.GetYaml()
	glog.SetLogFile("./log", "app.log")
	glog.SetCons(true)
}

func bootstrap() {
	glog.Errorf("config--->%+v", conf)
	server.NewServer().Start(fmt.Sprintf(":%d", conf.HttpConfig.Server.Port))
}

func main() {
	bootstrap()
}
