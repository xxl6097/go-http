package main

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/api"
	"github.com/xxl6097/go-http/api/test"
	"github.com/xxl6097/go-http/server"
	"github.com/xxl6097/go-http/server/route"
)

func init() {
	route.RouterUtil.SetApiPath("/v1/api")
	glog.SetLogFile("./log", "app.log")
	glog.SetCons(true)
}

func bootstrap() {
	api.GetApi().Add(test.NewRoute(test.NewController()))
	server.NewServer().Start(":9080")
}

func main() {
	bootstrap()
}
