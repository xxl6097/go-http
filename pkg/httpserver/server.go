package httpserver

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-http/internal"
	"github.com/xxl6097/go-http/internal/middle"
	"github.com/xxl6097/go-http/pkg/ihttpserver"
	"net/http"
)

var RouterUtil = middle.RouterUtil
var TokenUtils = middle.TokenUtils

type httpserver struct {
	routes []ihttpserver.IRoute
	server *internal.Server
}

func New() *httpserver {
	this := httpserver{
		routes: make([]ihttpserver.IRoute, 0),
	}
	this.server = internal.NewServer()

	return &this
}

func (this *httpserver) AddRoute(routes ...ihttpserver.IRoute) *httpserver {
	for _, route := range routes {
		this.routes = append(this.routes, route)
	}
	return this
}
func (this *httpserver) RouterFunc(fn func(router *mux.Router)) *httpserver {
	if fn != nil {
		fn(this.server.GetRouter())
	}
	return this
}

func (this *httpserver) register(router *mux.Router) {
	for _, route := range this.routes {
		route.Setup(router)
	}
}

func (this *httpserver) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) *httpserver {
	this.server.HandleFunc(pattern, handler)
	return this
}
func (this *httpserver) CORSMethodMiddleware() *httpserver {
	this.server.CORSMethodMiddleware()
	return this
}

func (this *httpserver) Use(mwf ...mux.MiddlewareFunc) *httpserver {
	this.server.Use(mwf...)
	return this
}

func (this *httpserver) Done(port int) {
	this.register(this.server.GetRouter())
	this.server.Start(fmt.Sprintf(":%d", port))
}

func (this *httpserver) DoneTSL(port int, cert tls.Certificate) {
	this.register(this.server.GetRouter())
	this.server.StartTSL(fmt.Sprintf(":%d", port), cert)
}
