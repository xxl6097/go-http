package test

import (
	"github.com/gorilla/mux"
	"go-http/server/inter"
	"go-http/server/util"
	"net/http"
)

type CardRoute struct {
	controller *TestController
}

func NewRoute(ctl *TestController) inter.IRoute {
	opt := &CardRoute{
		controller: ctl,
	}
	return opt
}

//func NewRoute(router *mux.Router, ctl *TestController) *CardRoute {
//	return &CardRoute{
//		router:     router,
//		controller: ctl,
//	}
//}

func (this *CardRoute) Setup(router *mux.Router) {
	util.RouterUtil.AddHandleFunc(router, util.ApiModel{
		Method: http.MethodPost,
		Path:   "/mqtt/auth",
		Fun:    this.controller.Auth,
		NoAuth: true,
	})
	util.RouterUtil.AddHandleFunc(router, util.ApiModel{
		Method: http.MethodPost,
		Path:   "/mqtt/post",
		Fun:    this.controller.Post,
		NoAuth: true,
	})
	util.RouterUtil.AddHandleFunc(router, util.ApiModel{
		Method: http.MethodGet,
		Path:   "/mqtt/test",
		Fun:    this.controller.Test,
		NoAuth: true,
	})
}
