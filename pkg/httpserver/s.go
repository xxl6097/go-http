package httpserver

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-http/internal/middle"
	"github.com/xxl6097/go-http/pkg/ihttpserver"
	"github.com/xxl6097/go-http/pkg/util"
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	router     *mux.Router
	server     *http.Server
	shutdownWG sync.WaitGroup
	routes     []ihttpserver.IRoute
}

func New() *Server {
	router := mux.NewRouter()
	router.Use(middle.EnableCors, middle.HandleOptions, middle.AuthMiddleware)
	this := &Server{
		router: router,
		routes: []ihttpserver.IRoute{},
	}
	return this
}
func (this *Server) CORSMethodMiddleware() *Server {
	this.router.Use(mux.CORSMethodMiddleware(this.router))
	return this
}
func (this *Server) AddRoute(routes ...ihttpserver.IRoute) *Server {
	for _, route := range routes {
		this.routes = append(this.routes, route)
	}
	return this
}

func (this *Server) AddApi(api func(router *mux.Router)) *Server {
	if api == nil {
		return this
	}
	api(this.router)
	return this
}
func (this *Server) Handle(pattern string, handler http.Handler) *Server {
	this.router.Handle(pattern, handler)
	return this
}
func BasicAuth(router *mux.Router, username, password string) {
	if username == "" || password == "" || router == nil {
		return
	}
	router.Use(middle.NewHTTPAuthMiddleware(username, password).SetAuthFailDelay(200 * time.Millisecond).Middleware)
}
func (this *Server) BasicAuth(username, password string) *Server {
	BasicAuth(this.router, username, password)
	return this
}
func (this *Server) BasicAuthRouter(router *mux.Router, username, password string) *Server {
	router.Use(middle.NewHTTPAuthMiddleware(username, password).SetAuthFailDelay(200 * time.Millisecond).Middleware)
	return this
}
func (this *Server) RouterFunc(fn func(router *mux.Router)) *Server {
	if fn != nil {
		fn(this.router)
	}
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

// Start 异步启动 HTTP 服务
func (this *Server) Start(address string) {
	this.shutdownWG.Add(1)
	s := &http.Server{
		Addr:    address,
		Handler: this.router,
	}
	this.server = s
	glog.Debug(fmt.Sprintf("http://%s%s", util.GetHostIp(), address))
	go func() {
		defer this.shutdownWG.Done()
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed: %v", err)
		}
	}()
}

func (this *Server) StartTSL(address string, cert tls.Certificate) {
	this.shutdownWG.Add(1)
	server := &http.Server{
		Addr:      address,
		Handler:   this.router,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
	}
	this.server = server
	ln, err := tls.Listen("tcp", address, server.TLSConfig)
	if err != nil {
		glog.Fatal("server listen err:", err)
		return
	}
	glog.Debug(fmt.Sprintf("https://%s%s", util.GetHostIp(), address))
	go func() {
		defer this.shutdownWG.Done()
		if e := server.Serve(ln); e != nil && !errors.Is(e, http.ErrServerClosed) {
			log.Fatalf("Server failed: %v", e)
		}
	}()
}

// Stop 安全关闭服务（含超时控制）
func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}
	s.shutdownWG.Wait()
	log.Println("Server stopped")
}

func (s *Server) Wait() {
	s.shutdownWG.Wait()
	log.Println("Server Wait")
}
func (this *Server) setup() {
	for _, route := range this.routes {
		route.Setup(this.router)
	}
}

func (this *Server) Done(port int) *Server {
	this.setup()
	this.Start(fmt.Sprintf(":%d", port))
	return this
}

func (this *Server) DoneTSL(port int, cert tls.Certificate) {
	this.setup()
	this.StartTSL(fmt.Sprintf(":%d", port), cert)
	defer this.Stop()
}
