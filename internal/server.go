package internal

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-http/internal/middle"
	"github.com/xxl6097/go-http/pkg/util"
	"net"
	"net/http"
)

//var (
//	httpServerReadTimeout  = 60 * time.Second
//	httpServerWriteTimeout = 60 * time.Second
//)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	this := &Server{
		router: mux.NewRouter(),
	}
	this.initRouter()
	return this
}

func (this *Server) initRouter() {
	//顺序，后面最先调用
	this.router.Use(middle.EnableCors, middle.HandleOptions, middle.AuthMiddleware)
	//this.router.Use(mux.CORSMethodMiddleware(this.router))
	//api.GetApi().Setup(this.router)
}

func (this *Server) CORSMethodMiddleware() *Server {
	this.router.Use(mux.CORSMethodMiddleware(this.router))
	return this
}
func (this *Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) *Server {
	this.router.HandleFunc(pattern, handler)
	return this
}

func (this *Server) GetRouter() *mux.Router {
	return this.router
}

func (this *Server) Use(mwf ...mux.MiddlewareFunc) *Server {
	this.router.Use(mwf...)
	return this
}

func (this *Server) Start(address string) {
	server := &http.Server{
		Addr:    address,
		Handler: this.router,
		//ReadTimeout:  httpServerReadTimeout,
		//WriteTimeout: httpServerWriteTimeout,
	}
	ln, err := net.Listen("tcp", address)
	if err != nil {
		glog.Fatal("server listen err:", err)
		return
	}
	glog.Debug(fmt.Sprintf("http://%s%s", util.GetHostIp(), address))
	_ = server.Serve(ln)
}

func (this *Server) StartTSL(address string, cert tls.Certificate) {
	server := &http.Server{
		Addr:    address,
		Handler: this.router,
		//ReadTimeout:  httpServerReadTimeout,
		//WriteTimeout: httpServerWriteTimeout,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
	}
	ln, err := tls.Listen("tcp", address, server.TLSConfig)
	if err != nil {
		glog.Fatal("server listen err:", err)
		return
	}
	glog.Debug(fmt.Sprintf("https://%s%s", util.GetHostIp(), address))
	_ = server.Serve(ln)
}
