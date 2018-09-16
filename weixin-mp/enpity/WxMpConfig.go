package enpity

import (
	"gopkg.in/go-playground/validator.v9"
)


const (
	secrst		=		"64ebc5c45e7531e88e9466ffb8435ff0"
	appid		=		"wx797295c89a602103"
)

//
type MpConfig struct {
	AppId					string		`validate:"gt=0"`
	Secret					string		`validate:"gt=0"`
	Token					string		`validate:"gt=0"`
	AesKey					string		`validate:"gt=0"`

	AccessToken 			string
	AccessTokenExpiresTime 	int64
	IsExpire				bool
	JsapiTicket				string
	JsapiTicketExpiresTime	int64
	OAuthToken				WxOAuthAccessToken
}

// 结构体参数验证
func Validator (config *MpConfig) error {
	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		return err
	}
	return nil
}
