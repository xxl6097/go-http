package middle

import (
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
		next.ServeHTTP(w, r)
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
