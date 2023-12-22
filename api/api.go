package api

import (
	"github.com/gorilla/mux"
	"go-http/api/test"
	"go-http/server/inter"
	"sync"
)

var instance *apiSingleton
var lock = &sync.Mutex{}

type apiSingleton struct {
	business []inter.IRoute
}

func GetApi() *apiSingleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &apiSingleton{
				business: make([]inter.IRoute, 0),
			}
		}
	}
	return instance
}

func (this *apiSingleton) add(routes ...inter.IRoute) {
	for _, route := range routes {
		this.business = append(this.business, route)
	}
}

func (this *apiSingleton) init() {
	this.add(test.NewRoute(test.NewController()))
}

func (this *apiSingleton) Setup(router *mux.Router) {
	this.init()
	for _, route := range this.business {
		route.Setup(router)
	}
}
