package main

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/api"
	"github.com/xxl6097/go-http/api/test"
	"github.com/xxl6097/go-http/server"
	"github.com/xxl6097/go-http/server/route"
	"github.com/xxl6097/go-http/server/util"
)

type HttpServerConfig struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	ApiPath string `yaml:"apipath"`
}

type HttpConfig struct {
	Server HttpServerConfig `yaml:"server"`
}
type Home struct {
	Home       string     `yaml:"home"`
	HttpConfig HttpConfig `yaml:"http"`
}

var conf = &Home{}

func init() {
	util.ParseYaml(conf)
	route.RouterUtil.SetApiPath(conf.HttpConfig.Server.ApiPath)
	fmt.Println(conf)
	glog.SetLogFile("./log", "app.log")
	glog.SetCons(true)
}

func bootstrap() {
	glog.Errorf("config--->%+v", conf)
	api.GetApi().Add(test.NewRoute(test.NewController()))
	server.NewServer().Start(fmt.Sprintf(":%d", conf.HttpConfig.Server.Port))
}

func main() {
	bootstrap()
}
