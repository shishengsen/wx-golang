package service

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
	"sync"
	"weixin-golang/weixin-mp/enpity"
	"weixin-golang/weixin-common/log"
	"weixin-golang/weixin-common/http"
	"weixin-golang/weixin-common/crypto"
)

const (
	access_token_url		=		"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

type WeChat struct {
	Cfg			*enpity.MpConfig
}

var once sync.Once
var weChat *WeChat

// 获取配置对象
func GetWeChat() WeChat {
	return *weChat
}

// 确保只初始化一次 MpConfig
func WxMpConfigStoreInMem(cfg *enpity.MpConfig) WeChat {
	once.Do(func(){
		cfg.IsExpire = false
		log := log.GetLogger()
		err := enpity.Validator(cfg)
		if err == nil {
			weChat.Cfg = cfg
		} else {
			log.Error(err)
		}
	})
	return *weChat
}

// 将微信配置信息存储到redis
func (w *WeChat)WxMpConfigStoreInRedis(cfg *enpity.MpConfig) {
	
}

// 将token信息存储在内存中
func (w *WeChat)wxOAuthTokenStoreInMem(oauth enpity.WxOAuthAccessToken) {
	lock := sync.NewCond(new(sync.Mutex))
	lock.L.Lock()
	w.Cfg.OAuthToken = oauth
	lock.L.Unlock()
}

// 将token信息存储到redis中
func (w *WeChat)wxOAuthTokenStoreInRedis(oauth enpity.WxOAuthAccessToken) {

}

// 签名验证
func (w *WeChat)CheckSignature(cfg enpity.MpConfig, signature string, timestamp string, nonce string) bool {
	_signature := crypto.Sha1(cfg.Token, timestamp, nonce)
	if _signature == signature {
		return true
	}
	panic("Error{检验signature}失败")
}

// 刷新accessToken信息
func (w *WeChat)UpdateAccessToken() {
	lock := sync.NewCond(new(sync.Mutex))
	lock.L.Lock()
	if isExpires() {
		tokenMap := w.refreshToken(w.Cfg)
		w.Cfg.AccessTokenExpiresTime = tokenMap["expires_in"].(int64) + time.Now().Unix()
		w.Cfg.AccessToken = tokenMap["access_token"].(string)
	}
	lock.L.Unlock()
}

// 获取accessToken信息
func (w *WeChat)GetAccessToken() string {
	if isExpires() {

	}
	return (*w.Cfg).AccessToken
}

// 内部调用刷新accessToken的微信api接口，此处是真正实现accessToken刷新的方法
func (w *WeChat)refreshToken(cfg *enpity.MpConfig) map[string]interface{} {
	requestUrl := fmt.Sprintf(access_token_url, cfg.AppId, cfg.Secret)
	msg, _ := http.Get(requestUrl)
	var f interface{}
	json.Unmarshal(msg, &f)
	m := f.(map[string]interface{})
	return m
}

// 解析微信返回的xml数据
func (w *WeChat)wxMpSubscribeMsgService(buf []byte) enpity.WxMessage {
	var msg enpity.WxMessage
	xml.Unmarshal(buf, &msg)
	return msg
}

func (w *WeChat)wxMpMsgAseDesc() string {
	return ""
}

// 检查微信功能调用的accessToken是否过期（注意，这里的accessToken不是获取用户信息的OAuth accessToken）
func isExpires() bool {
	return GetWeChat().Cfg.AccessTokenExpiresTime < time.Now().Unix()
}
