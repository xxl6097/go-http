package server

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/gologview/go/glogweb"
	logutil "github.com/xxl6097/gologview/go/util"
	"go-http/api"
	middle2 "go-http/server/middle"
	"go-http/server/util"
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
	glogweb.GetLogApi().HandlerLogView(this.router.NewRoute().Subrouter(), "admin", "het002402")
	//顺序，后面最先调用
	this.router.Use(middle2.EnableCors, middle2.HandleOptions, middle2.AuthMiddleware)
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

	ip := logutil.GetHostIp()
	glog.Errorf("api host http://%s%s\n", ip, server.Addr)
	glog.Errorf("log addr http://%s%s/logview/\n", ip, server.Addr)
	//this.router.Use(mux.CORSMethodMiddleware(this.router))

	glog.Debug(util.NotLoginUri)
	glog.Debug(util.NotLoginUriByPrefix)
	_ = server.Serve(ln)
}