package static

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-http/server/inter"
	"github.com/xxl6097/go-http/server/route"
	"net/http"
	"os"
)

var (
	static_prefix string = "/files/"
)

type static struct {
}

func (this *static) Setup(_router *mux.Router) {
	router := _router.NewRoute().Subrouter()
	baseDir, _ := os.Getwd()
	route.RouterUtil.AddNoAuthPrefix("files")
	//route.RouterUtil.AddNoAuthPrefix("/files")
	//route.RouterUtil.AddNoAuthPrefix("/files/")
	// 创建一个静态文件服务器，处理 "/static/" 路径
	fileServer := http.FileServer(http.Dir(baseDir))
	handle := http.StripPrefix(static_prefix, fileServer)
	router.PathPrefix(static_prefix).Handler(handle)
}

func NewRoute() inter.IRoute {
	opt := &static{}
	return opt
}
