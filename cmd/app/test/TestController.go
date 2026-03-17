package test

import (
	"github.com/xxl6097/glog/pkg/z"
	"github.com/xxl6097/go-http/pkg/util"
	"go.uber.org/zap"

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
	z.L().Warn("Test", zap.Any("r", r))
	Respond(w, Ignore(false))
}

func (this *TestController) Post(w http.ResponseWriter, r *http.Request) {
	req, _ := util.GetReqMapData(w, r)
	if req != nil {
		z.L().Warn("resp", zap.Any("req", req))
	}

	Respond(w, Ignore(false))
}

func (this *TestController) Auth(w http.ResponseWriter, r *http.Request) {
	req, _ := util.GetReqMapData(w, r)
	z.L().Warn("Auth", zap.Any("req", req))
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

func (this *TestController) Frp(w http.ResponseWriter, r *http.Request) {
	req, _ := util.GetReqMapData(w, r)
	z.L().Warn("Frp", zap.Any("req", req))
	data := map[string]interface{}{
		"reject":   false,
		"unchange": true,
	}
	Respond(w, data)
}
