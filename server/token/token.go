package token

var TokenUtils = tokenUtil{}

type TokenHandler func(string) (bool, map[string]interface{})

//type TokenHandler interface {
//	CheckToken(token string) (bool, map[string]interface{})
//}

type TokenModel struct {
	AccessToken    string `json:"accessToken"`
	AccessUuid     string `json:"accessUuid"`
	AccessExpires  int64  `json:"accessExpires"`
	RefreshToken   string `json:"refreshToken"`
	RefreshUuid    string `json:"refreshUuid"`
	RefreshExpires int64  `json:"refreshExpires"`
}

// tokenUtil token操作工具类
type tokenUtil struct {
	callback TokenHandler
}

func (this *tokenUtil) SetTokenCallBack(callback TokenHandler) {
	this.callback = callback
}

func (this *tokenUtil) CheckToken(token string, f func(bool, map[string]interface{})) {
	if this.callback != nil {
		ok, maps := this.callback(token)
		if f != nil {
			f(ok, maps)
		}
	} else {
		if f != nil {
			f(true, nil)
		}
	}
}
