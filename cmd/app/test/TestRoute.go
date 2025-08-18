package test

import (
	"github.com/gorilla/mux"
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
	router.HandleFunc("/mqtt/auth", this.controller.Auth).Methods(http.MethodPost)
	router.HandleFunc("/mqtt/post", this.controller.Post).Methods(http.MethodPost)
	router.HandleFunc("/mqtt/test", this.controller.Test).Methods(http.MethodGet)
	router.HandleFunc("/prod-api/vue-admin-template/user/info", this.controller.Test).Methods(http.MethodGet)
	router.HandleFunc("/frp", this.controller.Frp).Methods(http.MethodPost)
}
