package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xxl6097/glog/glog"
	"net"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

var (
	Success = map[string]interface{}{"code": 0, "msg": "成功"}
	Failed  = map[string]interface{}{"code": 300, "msg": "失败"}

	ParamsValidError    = map[string]interface{}{"code": 310, "msg": "参数校验错误"}
	ParamsTypeError     = map[string]interface{}{"code": 311, "msg": "参数类型错误"}
	RequestMethodError  = map[string]interface{}{"code": 312, "msg": "请求方法错误"}
	AssertArgumentError = map[string]interface{}{"code": 313, "msg": "断言参数错误"}

	LoginAccountError = map[string]interface{}{"code": 330, "msg": "登录账号或密码错误"}
	LoginDisableError = map[string]interface{}{"code": 331, "msg": "登录账号已被禁用了"}
	TokenEmpty        = map[string]interface{}{"code": 332, "msg": "token参数为空"}
	TokenInvalid      = map[string]interface{}{"code": 333, "msg": "token参数无效"}
	TokenExpires      = map[string]interface{}{"code": 334, "msg": "token参数过期"}

	NoPermission    = map[string]interface{}{"code": 403, "msg": "无相关权限"}
	Request404Error = map[string]interface{}{"code": 404, "msg": "请求接口不存在"}
	Request405Error = map[string]interface{}{"code": 405, "msg": "请求方法不允许"}

	SystemError = map[string]interface{}{"code": 500, "msg": "系统错误"}
)

func MessageError(msg string) map[string]interface{} {
	return map[string]interface{}{"code": -1, "msg": msg}
}

func MessageSucess(data interface{}) map[string]interface{} {
	return map[string]interface{}{"code": 0, "msg": "sucess", "data": data}
}
func MessageSucessOnMsg(msg string) map[string]interface{} {
	return map[string]interface{}{"code": 0, "msg": msg}
}
func MessageWithData(code int, msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{"code": code, "msg": msg, "data": data}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	//w.Header().Add("Content-Type", "application/json")
	//w.Header().Add("Access-Control-Allow-Origin", "*")
	//w.Header().Add("Access-Control-Allow-Headers", "content-type,Authorization")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	//w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
	//buf, _ := json.Marshal(data)
	//glog.Info(string(buf))
	if json.NewEncoder(w).Encode(data) != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetRequestParam(r *http.Request, key string) string {
	vals := r.URL.Query()[key]
	glog.Info("GetRequestParam, vals: ", vals)
	if vals == nil || len(vals) == 0 {
		return ""
	}
	return vals[0]
}

func GetReqData[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	var t T
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		glog.Error("DecodeBody error", err)
		return nil, err
	}
	return &t, nil
}

func GetReqMapData(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	model, err := GetReqData[map[string]interface{}](w, r)
	if model == nil {
		return nil, err
	}
	delete(*model, "ID")
	delete(*model, "CreatedAt")
	delete(*model, "UpdatedAt")
	delete(*model, "DeletedAt")
	for key, value := range *model {
		if value == nil { //null 提出
			delete(*model, key)
		}
	}
	return *model, err
}

// Contains 判断src是否包含elem元素
func Contains(src interface{}, elem interface{}) bool {
	srcArr := reflect.ValueOf(src)
	if srcArr.Kind() == reflect.Ptr {
		srcArr = srcArr.Elem()
	}
	if srcArr.Kind() == reflect.Slice {
		for i := 0; i < srcArr.Len(); i++ {
			if srcArr.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

func StartWith(prefix string, str string) bool {
	var headerData = []byte(prefix)
	if bytes.HasPrefix([]byte(str), headerData) {
		return true
	}
	return false
}

func StartWithByArr(str string, prefixs []string) bool {
	if prefixs == nil {
		return false
	}
	for _, prefix := range prefixs {
		has := strings.HasPrefix(str, prefix)
		if has {
			return has
		}
	}
	return false
}

func in(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}

//func CreateAuthHandler(fun func(http.ResponseWriter, *http.Request)) http.Handler {
//	return middle.EnableCors(middle.HandleOptions(middle.AuthMiddleware(http.HandlerFunc(fun))))
//}
//func CreateNoAuthHandler(fun func(http.ResponseWriter, *http.Request)) http.Handler {
//	return middle.EnableCors(middle.HandleOptions(http.HandlerFunc(fun)))
//}

func GetHostIp() string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	//var ips []net.IP
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.IsPrivate() {
			if ipNet.IP.To4() != nil {
				//ip = ipNet.IP.String()
				//break
				ip := ipNet.IP.To4()
				//fmt.Println(ip[0])
				switch ip[0] {
				case 10:
					return ipNet.IP.String()
				case 192:
					return ipNet.IP.String()
				}
			}
		}
	}
	//fmt.Println(ips)
	return ""
}
