package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var RouterUtil = routerUtil{}

// 免登录验证
var NotLoginUri = []string{}
var NotLoginUriByPrefix = []string{"logview", "fserver"}
var Apis = []string{}
var FServers = []string{}

type ApiModel struct {
	Fun            func(http.ResponseWriter, *http.Request)
	Method         string
	Path           string
	NoAuth         bool
	NoAuthByPrefix bool
}

// redisUtil Redis操作工具类
type routerUtil struct {
	apipath string
}

func (this *routerUtil) SetApiPath(path string) {
	this.apipath = path
}

func (this *routerUtil) AddHandleFunc(router *mux.Router, models ...ApiModel) {
	for _, model := range models { //config.Get().GetConfig().HttpConfig.Server.ApiPath
		apipth := this.apipath + model.Path
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
	for _, model := range models { //config.Get().GetConfig().HttpConfig.Server.ApiPath
		apipth := this.apipath + model.Path
		auths := strings.ReplaceAll(strings.Replace(apipth, "/", "", 1), "/", ":")
		if model.NoAuth {
			NotLoginUri = append(NotLoginUri, auths)
		}
		if model.NoAuthByPrefix {
			NotLoginUriByPrefix = append(NotLoginUriByPrefix, auths)
		}
		//glog.Warn(apipth)
		FServers = append(FServers, apipth)
		router.PathPrefix(apipth).Handler(http.StripPrefix(apipth, http.FileServer(http.Dir("./files"))))
	}
}

func (this *routerUtil) AddNoAuth(url string) {
	if url != "" {
		NotLoginUri = append(NotLoginUri, url)
	}
}

func (this *routerUtil) AddNoAuthPrefix(prefix string) {
	if prefix != "" {
		NotLoginUriByPrefix = append(NotLoginUriByPrefix, prefix)
	}
}
