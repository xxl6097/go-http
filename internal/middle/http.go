package middle

import (
	"compress/gzip"
	"crypto/subtle"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type HTTPAuthMiddleware struct {
	user          string
	passwd        string
	authcodes     []string
	authFunc      func(r *http.Request) bool
	authFailDelay time.Duration
}

func NewHTTPAuthMiddleware(user, passwd string) *HTTPAuthMiddleware {
	return &HTTPAuthMiddleware{
		user:      user,
		passwd:    passwd,
		authcodes: make([]string, 0),
	}
}

func (authMid *HTTPAuthMiddleware) AddAuthCode(codes ...string) *HTTPAuthMiddleware {
	if codes != nil {
		for _, code := range codes {
			authMid.authcodes = append(authMid.authcodes, code)
		}
	}
	return authMid
}

func (authMid *HTTPAuthMiddleware) SetAuthFailDelay(delay time.Duration) *HTTPAuthMiddleware {
	authMid.authFailDelay = delay
	return authMid
}

func (authMid *HTTPAuthMiddleware) AuthFunc(fn func(r *http.Request) bool) *HTTPAuthMiddleware {
	authMid.authFunc = fn
	return authMid
}
func (authMid *HTTPAuthMiddleware) checkBasic(next http.Handler, w http.ResponseWriter, r *http.Request) bool {
	//autoCode := r.URL.Query().Get("auth_code")
	//if autoCode == "" {
	//	query, err := url.Parse(r.Referer())
	//	if err == nil && query != nil {
	//		if query.Query().Has("auth_code") {
	//			autoCode = query.Query().Get("auth_code")
	//		}
	//	}
	//}
	//glog.Infof("Auth: %s", autoCode)
	//glog.Infof("RequestURI: %s", r.RequestURI)
	//glog.Infof("Referer: %s", r.Referer())
	if authMid.authFunc != nil {
		ok := authMid.authFunc(r)
		if ok {
			next.ServeHTTP(w, r)
			return true
		}
	}
	reqUser, reqPasswd, hasAuth := r.BasicAuth()
	if authMid.user == "" && authMid.passwd == "" {
		next.ServeHTTP(w, r)
		return true
	} else if hasAuth {
		//if autoCode != "" {
		//	auth := reqUser + ":" + reqPasswd
		//	code := base64.StdEncoding.EncodeToString([]byte(auth))
		//	if ConstantTimeEqString(code, autoCode) {
		//		next.ServeHTTP(w, r)
		//		return true
		//	}
		//} else
		if ConstantTimeEqString(reqUser, authMid.user) && ConstantTimeEqString(reqPasswd, authMid.passwd) {
			next.ServeHTTP(w, r)
			return true
		}
	}
	return false
}

func (authMid *HTTPAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !authMid.checkBasic(next, w, r) {
			if authMid.authFailDelay > 0 {
				time.Sleep(authMid.authFailDelay)
			}
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			//http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			http.Error(w, fmt.Sprintf("%s，未授权，请输入basic授权信息", http.StatusText(http.StatusUnauthorized)), http.StatusUnauthorized)
		}
	})
}

type HTTPGzipWrapper struct {
	h http.Handler
}

func (gw *HTTPGzipWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		gw.h.ServeHTTP(w, r)
		return
	}
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()
	gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
	gw.h.ServeHTTP(gzr, r)
}

func MakeHTTPGzipHandler(h http.Handler) http.Handler {
	return &HTTPGzipWrapper{
		h: h,
	}
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func ConstantTimeEqString(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
