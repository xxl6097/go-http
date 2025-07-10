package main

import (
	"github.com/xxl6097/glog/glog"
	"github.com/xxl6097/go-http/cmd/app/test"
	"github.com/xxl6097/go-http/internal/middle"
	"github.com/xxl6097/go-http/pkg/httpserver"
)

//func init() {
//	//route.RouterUtil.SetApiPath("/v1/api")
//	glog.SetLogFile("./log", "app.log")
//	glog.SetCons(true)
//}

//func initSse(router *mux.Router) {
//	isse := sse.New().
//		InvalidateFun(func(r *http.Request) (string, error) {
//			uuid := r.URL.Query().Get("uuid")
//			//session := v5x.GetInstance().GetUser(uuid)
//			//if session == nil {
//			//	glog.Error("sse用户连接不合法", uuid)
//			//	return time.Now().Format("20060102150405.999999999"), fmt.Errorf("sse用户连接不合法 %s", uuid)
//			//}
//			//r.Header.Set("Sse-Event-GroupID", session.UserKey)
//			return uuid, nil
//		}).
//		Register(func(server iface.ISseServer, client *iface.Client) {
//			//server.Stream("国家主席习近平在陕西省西安市", time.Second)
//		}).
//		Done()
//	router.HandleFunc("/sse/stream", isse.Handler())
//}

func bootstrap() {
	middle.TokenUtils.Callback(func(s string) (bool, map[string]interface{}) {
		glog.Println("Callback", s)
		return true, nil
	})
	httpserver.New().
		Use(middle.NewHTTPAuthMiddleware("admin", "het002402").Middleware).
		AddRoute(test.NewRoute(test.NewController())).
		AddRoute(test.NewWsRoute()).
		//RouterFunc(initSse).
		Done(8080)
}

func main() {
	bootstrap()
}
