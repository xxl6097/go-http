package inter

import "github.com/gorilla/mux"

type IRoute interface {
	Setup(router *mux.Router)
}
