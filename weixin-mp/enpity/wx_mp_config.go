package enpity

import (
	"gopkg.in/go-playground/validator.v9"
	"sync"
)

const (
	/**
	测试所用
	 */
	secrst = "64ebc5c45e7531e88e9466ffb8435ff0"
	/**
	测试所用
	 */
	appid  = "wx797295c89a602103"
)

//
type MpConfig struct {
	AppId  						string `validate:"gt=0"`
	Secret 						string `validate:"gt=0"`
	Token  						string `validate:"gt=0"`
	AesKey 						string `validate:"gt=0"`

	AccessToken            		string
	AccessTokenExpiresTime 		int64
	IsExpire               		bool
	OAuthToken             		*WxOAuthAccessToken
	JsApiTicket					*WxJsTicket
	JsApiConfig					*WxJsConfig
	Lang						string

	AccessTokenLock				*sync.Cond			// 刷新微信授权码时需进行并发控制，防止过度刷新，需要整个config用一个锁
	JsapiTicketLock				*sync.Cond			// 刷新微信jssdk的ticket时需进行并发控制，防止过度刷新，需要整个config用一个锁

}

// 结构体参数验证
func Validator(config *MpConfig) error {
	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		return err
	}
	return nil
}
