package main

import (
	"github.com/xxl6097/go-glog/glog"
	test2 "github.com/xxl6097/go-http/cmd/app/test"
	"github.com/xxl6097/go-http/pkg/api"
	"github.com/xxl6097/go-http/server"
	"github.com/xxl6097/go-http/server/token"
)

func init() {
	//route.RouterUtil.SetApiPath("/v1/api")
	glog.SetLogFile("./log", "app.log")
	glog.SetCons(true)
}

func bootstrap() {
	token.TokenUtils.Callback(func(s string) (bool, map[string]interface{}) {
		glog.Println("Callback", s)
		return true, nil
	})
	api.GetApi().Add(test2.NewRoute(test2.NewController()))
	server.NewServer().Start(":8888")
}

func main() {
	bootstrap()
}
