package api

import (
	"github.com/gorilla/mux"
	"github.com/xxl6097/glog/pkg/z"
	"github.com/xxl6097/go-http/pkg/httpserver"
	"go.uber.org/zap"

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
	z.L().Info("apiHello", zap.String("Method", r.Method), zap.String("URL", r.URL.String()), zap.String("Proto", r.Proto))
	w.Write([]byte("hello world"))
}
