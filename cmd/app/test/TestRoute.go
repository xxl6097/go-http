package test

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-http/server/inter"
	"github.com/xxl6097/go-http/server/route"
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
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/mqtt/auth",
		Fun:    this.controller.Auth,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodPost,
		Path:   "/mqtt/post",
		Fun:    this.controller.Post,
		NoAuth: true,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodGet,
		Path:   "/mqtt/test",
		Fun:    this.controller.Test,
		NoAuth: false,
	})
	route.RouterUtil.AddHandleFunc(router, route.ApiModel{
		Method: http.MethodGet,
		Path:   "/prod-api/vue-admin-template/user/info",
		Fun:    this.controller.Test,
		NoAuth: false,
	})
}
