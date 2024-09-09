package main

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/api"
	"github.com/xxl6097/go-http/api/test"
	"github.com/xxl6097/go-http/server"
	"github.com/xxl6097/go-http/server/route"
	"github.com/xxl6097/go-http/server/token"
)

func init() {
	route.RouterUtil.SetApiPath("/v1/api")
	glog.SetLogFile("./log", "app.log")
	glog.SetCons(true)
}

func bootstrap() {
	token.TokenUtils.Callback(func(s string) (bool, map[string]interface{}) {
		glog.Println("Callback", s)
		return true, nil
	})
	api.GetApi().Add(test.NewRoute(test.NewController()))
	server.NewServer().Start(":8888")
}

func main() {
	bootstrap()
}
