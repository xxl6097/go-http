package util

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server/config"
	"net/http"
	"strings"
)

var RouterUtil = routerUtil{}

// 免登录验证
var NotLoginUri = []string{}
var NotLoginUriByPrefix = []string{"logview"}
var Apis = []string{}

type ApiModel struct {
	Fun            func(http.ResponseWriter, *http.Request)
	Method         string
	Path           string
	NoAuth         bool
	NoAuthByPrefix bool
}

// redisUtil Redis操作工具类
type routerUtil struct {
}

func (this *routerUtil) AddHandleFunc(router *mux.Router, models ...ApiModel) {
	for _, model := range models {
		apipth := config.Get().GetConfig().HttpConfig.Server.ApiPath + model.Path
		auths := strings.ReplaceAll(strings.Replace(apipth, "/", "", 1), "/", ":")
		if model.NoAuth {
			NotLoginUri = append(NotLoginUri, auths)
		}
		if model.NoAuthByPrefix {
			NotLoginUriByPrefix = append(NotLoginUriByPrefix, auths)
		}
		//glog.Warn(apipth)
		Apis = append(Apis, apipth)
		router.HandleFunc(apipth, model.Fun).Methods(model.Method, http.MethodOptions)
	}
}

func (this *routerUtil) AddFileServer(router *mux.Router, models ...ApiModel) {
	//this.router.PathPrefix("/v1/api/files/").Handler(http.StripPrefix("/v1/api/files/", http.FileServer(http.Dir("./files"))))
	for _, model := range models {
		apipth := config.Get().GetConfig().HttpConfig.Server.ApiPath + model.Path
		auths := strings.ReplaceAll(strings.Replace(apipth, "/", "", 1), "/", ":")
		if model.NoAuth {
			NotLoginUri = append(NotLoginUri, auths)
		}
		if model.NoAuthByPrefix {
			NotLoginUriByPrefix = append(NotLoginUriByPrefix, auths)
		}
		glog.Warn(apipth)
		router.PathPrefix(apipth).Handler(http.StripPrefix(apipth, http.FileServer(http.Dir("./files"))))
	}
}
