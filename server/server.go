package server

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/pkg/api"
	"github.com/xxl6097/go-http/server/middle"
	"github.com/xxl6097/go-http/server/route"
	"net"
	"net/http"
	"time"
)

var (
	httpServerReadTimeout  = 60 * time.Second
	httpServerWriteTimeout = 60 * time.Second
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	this := &Server{
		router: mux.NewRouter(),
	}
	this.initApi()
	return this
}

func (this *Server) initApi() {
	//顺序，后面最先调用
	this.router.Use(middle.EnableCors, middle.HandleOptions, middle.AuthMiddleware)
	this.router.Use(mux.CORSMethodMiddleware(this.router))
	api.GetApi().Setup(this.router)
}

func (this *Server) Start(address string) {
	server := &http.Server{
		Addr:         address,
		Handler:      this.router,
		ReadTimeout:  httpServerReadTimeout,
		WriteTimeout: httpServerWriteTimeout,
	}
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	glog.Debug(address)
	glog.Debug(route.NotLoginUri)
	glog.Debug(route.NotLoginUriByPrefix)
	_ = server.Serve(ln)
}
