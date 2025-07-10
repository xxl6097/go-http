package ihttpserver

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ApiModel struct {
	Fun            func(http.ResponseWriter, *http.Request)
	Method         string
	Path           string
	NoAuth         bool
	NoAuthByPrefix bool
}

type IRoute interface {
	Setup(router *mux.Router)
}
