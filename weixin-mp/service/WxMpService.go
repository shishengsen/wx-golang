package service

import (
	"encoding/json"
	"time"
	"sync"
	"weixin-golang/weixin-mp/enpity"
	"weixin-golang/weixin-common/log"
	"weixin-golang/weixin-common/http"
	"weixin-golang/weixin-common/crypto"
)

const (
	access_token_url		=		"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"
)

var mpConfig *enpity.MpConfig
var once sync.Once

// 确保只初始化一次 MpConfig
func WxMpConfigStoreInMem(cfg *enpity.MpConfig) {
	once.Do(func(){
		cfg.IsExpire = false
		log := log.GetLogger()
		err := enpity.Validator(cfg)
		if err == nil {
			mpConfig = cfg
		} else {
			log.Error(err)
		}
	})
	UpdateAccessToken()
}

//
func WxMpConfigStoreInRedis(cfg *enpity.MpConfig) {
	
}

//
func wxOAuthTokenStoreInMem(oauth enpity.WxOAuthAccessToken) {
	lock := sync.NewCond(new(sync.Mutex))
	lock.L.Lock()
	mpConfig.OAuthToken = oauth
	lock.L.Unlock()
}

func wxOAuthTokenStoreInRedis(oauth enpity.WxOAuthAccessToken) {

}

//
func GetMpConfig() enpity.MpConfig {
	return *mpConfig
}

func CheckSignature(cfg enpity.MpConfig, signature string, timestamp string, nonce string) bool {
	_signature := crypto.Sha1(cfg.Token, timestamp, nonce)
	if _signature == signature {
		return true
	}
	panic("Error{检验signature}失败")
}

//
func UpdateAccessToken() {
	lock := sync.NewCond(new(sync.Mutex))
	lock.L.Lock()
	tokenMap := refreshToken(mpConfig)
	mpConfig.AccessTokenExpiresTime = tokenMap["expires_in"].(int64) + time.Now().Unix()
	mpConfig.AccessToken = tokenMap["access_token"].(string)
	lock.L.Unlock()
}

//
func GetAccessToken() string {
	return (*mpConfig).AccessToken
}

//
func refreshToken(cfg *enpity.MpConfig) map[string]interface{} {
	requestUrl := access_token_url + "&appid=" + mpConfig.AppId + "&secret=" + mpConfig.Secret
	msg, _ := http.Get(requestUrl)
	var f interface{}
	json.Unmarshal(msg, &f)
	m := f.(map[string]interface{})
	return m
}
