package test

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server/util"
	"net/http"
	"strings"
)

type TestController struct {
}

func NewController() *TestController {
	return &TestController{}
}

func (this *TestController) Test(w http.ResponseWriter, r *http.Request) {
	//req := utils.GetReqMapData(w, r)
	//glog.Warn(req)
	glog.Warn("Test---->", r)
	Respond(w, Ignore(false))
}

func (this *TestController) Post(w http.ResponseWriter, r *http.Request) {
	req := util.GetReqMapData(w, r)
	if req != nil {
		glog.Warn("resp---->", req)
	}

	Respond(w, Ignore(false))
}

func (this *TestController) Auth(w http.ResponseWriter, r *http.Request) {
	req := util.GetReqMapData(w, r)
	glog.Warn(req)
	username := req["username"]
	if username == nil || username.(string) == "" {
		Respond(w, Deny(false))
		return
	}
	if strings.Compare("admin", username.(string)) == 0 {
		Respond(w, Allow(true))
		return
	}
	if strings.Compare("uuxia", username.(string)) == 0 {
		Respond(w, Allow(false))
		return
	}
	Respond(w, Ignore(false))
}
