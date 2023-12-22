package middle

import (
	"github.com/xxl6097/go-glog/glog"
	"go-http/server/token"
	"go-http/server/util"
	"net/http"
	"strings"
)

// AuthMiddleware 鉴权中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auths := strings.ReplaceAll(strings.Replace(r.URL.Path, "/", "", 1), "/", ":")
		if util.Contains(util.NotLoginUri, auths) {
			next.ServeHTTP(w, r)
			return
		}
		if util.StartWithByArr(auths, util.NotLoginUriByPrefix) {
			next.ServeHTTP(w, r)
			return
		}
		tk := r.Header.Get("accessToken")
		if tk == "" {
			//glog.Info("AuthMiddleware----2", tk)
			util.Respond(w, util.TokenEmpty)
			return
		}
		//next.ServeHTTP(w, r)
		token.TokenUtils.CheckToken(tk, func(ok bool, parms map[string]interface{}) {
			if ok {
				for k, v := range parms {
					r.Header.Set(k, v.(string))
				}
				next.ServeHTTP(w, r)
			} else {
				glog.Info(auths, http.StatusUnauthorized)
				w.WriteHeader(http.StatusUnauthorized)
			}
		})
		//isValidata, username, res := utils.TokenUtils.CheckToken(tk)
		//if isValidata {
		//	r.Header.Set("UserName", username)
		//	next.ServeHTTP(w, r)
		//} else {
		//	glog.Info(res)
		//	w.WriteHeader(http.StatusUnauthorized)
		//	//utils.Respond(w, res)
		//}
	})
}
