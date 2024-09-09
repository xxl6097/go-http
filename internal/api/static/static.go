package static

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server/inter"
	"github.com/xxl6097/go-http/server/route"
	"github.com/xxl6097/go-http/server/util"
	"net/http"
	"os"
)

var (
	static_prefix string = "/files/"
)

type static struct {
	username, password string
}

func (this *static) Setup(_router *mux.Router) {
	router := _router.NewRoute().Subrouter()
	if this.username != "" && this.password != "" {
		router.Use(util.NewHTTPAuthMiddleware(this.username, this.password).Middleware)
	}
	baseDir, _ := os.Getwd()
	route.RouterUtil.AddNoAuthPrefix("files")
	//route.RouterUtil.AddNoAuthPrefix("/files")
	//route.RouterUtil.AddNoAuthPrefix("/files/")
	// 创建一个静态文件服务器，处理 "/static/" 路径
	glog.Println("baseDir:", baseDir)
	fileServer := http.FileServer(http.Dir(baseDir))
	handle := http.StripPrefix(static_prefix, fileServer)
	router.PathPrefix(static_prefix).Handler(handle)
}

func NewRoute(username, password string) inter.IRoute {
	opt := &static{
		username: username,
		password: password,
	}
	return opt
}
