package config

import (
	"encoding/json"
	"weixin-golang/weixin-common/http"
	"weixin-golang/weixin-mp/enpity"
)

const (
	access_token		=		"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"
	secrst				=		"64ebc5c45e7531e88e9466ffb8435ff0"
	appid				=		"wx797295c89a602103"
)

//
func refreshToken(cfg *enpity.MpConfig) map[string]interface{} {
	requestUrl := access_token + "&appid=" + mpConfig.AppId + "&secret=" + mpConfig.Secret
	msg, _ := http.Get(requestUrl)
	var f interface{}
	json.Unmarshal(msg, &f)
	m := f.(map[string]interface{})
	return m
}
