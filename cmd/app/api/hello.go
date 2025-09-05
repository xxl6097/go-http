package api

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-http/pkg/httpserver"
	"net/http"
)

func NewHello() func(*mux.Router) {
	return func(router *mux.Router) {
		hRouter := router.NewRoute().Subrouter()
		//httpserver.BasicAuth(hRouter, "admin", "admin")
		httpserver.BasicAuth(hRouter, "admin", "admin", "het002402")
		hRouter.HandleFunc("/api/hello", apiHello)
	}
}
func apiHello(w http.ResponseWriter, r *http.Request) {
	glog.Printf("%s %s %s\n", r.Method, r.URL.String(), r.Proto)
	w.Write([]byte("hello world"))
}
