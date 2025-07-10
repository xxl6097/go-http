package test

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-http/pkg/httpserver"
	"github.com/xxl6097/go-http/pkg/ihttpserver"
	"net/http"
)

type CardRoute struct {
	controller *TestController
}

func NewRoute(ctl *TestController) ihttpserver.IRoute {
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
	httpserver.RouterUtil.AddHandleFunc(router, ihttpserver.ApiModel{
		Method: http.MethodPost,
		Path:   "/mqtt/auth",
		Fun:    this.controller.Auth,
		NoAuth: true,
	})
	httpserver.RouterUtil.AddHandleFunc(router, ihttpserver.ApiModel{
		Method: http.MethodPost,
		Path:   "/mqtt/post",
		Fun:    this.controller.Post,
		NoAuth: true,
	})
	httpserver.RouterUtil.AddHandleFunc(router, ihttpserver.ApiModel{
		Method: http.MethodGet,
		Path:   "/mqtt/test",
		Fun:    this.controller.Test,
		NoAuth: false,
	})
	httpserver.RouterUtil.AddHandleFunc(router, ihttpserver.ApiModel{
		Method: http.MethodGet,
		Path:   "/prod-api/vue-admin-template/user/info",
		Fun:    this.controller.Test,
		NoAuth: false,
	})
	httpserver.RouterUtil.AddHandleFunc(router, ihttpserver.ApiModel{
		Method: http.MethodPost,
		Path:   "/frp",
		Fun:    this.controller.Frp,
		NoAuth: false,
	})
}
