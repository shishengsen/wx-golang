package service

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"sync"
	"time"
	"wx-golang/weixin-common/crypto"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/utils"
	"wx-golang/weixin-mp/enpity"
)

const (
	/**

	 */
	access_token_url 		=		"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	/**

	 */
	clear_quota      		= 		"https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=%s"
	/**

	 */
	jsapi_ticket     		=		"https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	/**

	 */
	jsapi_signature  		= 		"jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s"
	/**
	长链接转短链接接口
	 */
	longurl_to_shorturl		=		"https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s"
)

type WeChat struct {
	cfg     		*enpity.MpConfig
	router				*MsgRouter
}

// 确保只初始化一次 MpConfig
func WxNewMpConfig(cfg *enpity.MpConfig) WeChat {
	cfg.IsExpire = false
	err := enpity.Validator(cfg)
	if err != nil {
		panic(err)
	}
	weChat := new(WeChat)
	cfg.AccessTokenLock = sync.NewCond(new(sync.Mutex))
	cfg.JsapiTicketLock = sync.NewCond(new(sync.Mutex))
	weChat.cfg = cfg
	return *weChat
}

// 将微信配置信息存储到redis
func (w *WeChat) WxMpConfigStoreInRedis(cfg *enpity.MpConfig) {

}

// 将token信息存储在内存中
func (w *WeChat) wxOAuthTokenStoreInMem(oauth *enpity.WxOAuthAccessToken) {
	w.cfg.AccessTokenLock.L.Lock()
	w.cfg.OAuthToken = oauth
	w.cfg.AccessTokenLock.L.Unlock()
}

// 将token信息存储到redis中
func (w *WeChat) wxOAuthTokenStoreInRedis(oauth enpity.WxOAuthAccessToken) {

}

// 签名验证
func (w *WeChat) WxCheckSignature(cfg enpity.MpConfig, signature string, timestamp string, nonce string) bool {
	_signature := crypto.Sha1(cfg.Token, timestamp, nonce)
	if _signature == signature {
		return true
	}
	panic("Error{检验signature}失败")
}

// 刷新accessToken信息，token并发控制，防止过度刷新token信息导致access_token调用次数用完
func (w *WeChat) WxUpdateAccessToken() {
	w.cfg.AccessTokenLock.L.Lock()
	if isExpires(*w) {
		tokenMap := refreshToken(w.cfg, *w)
		w.cfg.AccessTokenExpiresTime = tokenMap["expires_in"].(int64) + time.Now().Unix()
		w.cfg.AccessToken = tokenMap["access_token"].(string)
	}
	w.cfg.AccessTokenLock.L.Unlock()
}

// 获取accessToken信息
func (w *WeChat) WxGetAccessToken() string {
	if isExpires(*w) {
		w.WxUpdateAccessToken()
	}
	return (*w.cfg).AccessToken
}

// 公众号调用或第三方平台帮公众号调用对公众号的所有api调用（包括第三方帮其调用）次数进行清零：
func (w *WeChat) WxApiClearQuota() map[string]interface{} {
	reqUrl := fmt.Sprintf(clear_quota, w.WxGetAccessToken())
	reqBody := map[string]string{
		"appid": w.cfg.AppId,
	}
	msg := http.Post(reqUrl, string(utils.Interface2byte(reqBody)))
	var respBody map[string]interface{}
	json.Unmarshal(msg, &reqBody)
	return respBody
}

// 解析微信返回的xml数据
func (w *WeChat) WxMpSubscribeMsgService(buf []byte) enpity.WxMessage {
	var msg enpity.WxMessage
	xml.Unmarshal(buf, &msg)
	return msg
}

// 获取微信js的ticket
func (w *WeChat) WxGetWxJsApiTicket(forceRefresh bool) enpity.WxJsTicket {
	w.cfg.JsapiTicketLock.L.Lock()
	if forceRefresh {
		getWxJsapiTicket(w.cfg)
	}
	if !forceRefresh && w.cfg.JsApiTicket.IsExpires() {
		getWxJsapiTicket(w.cfg)
	}
	w.cfg.JsapiTicketLock.L.Unlock()
	return *w.cfg.JsApiTicket
}

// 返回微信jssdk使用所需要的信息
func (w *WeChat) WxCreateJsapiSignature(url string) enpity.WxJsConfig {
	appid := w.cfg.AppId
	jsticket := w.cfg.JsApiTicket.Ticket
	timestamp := time.Now().Unix()
	noncestr := utils.RandomStr()
	signature := crypto.Sha1WithAmple(fmt.Sprintf(jsapi_signature, jsticket, noncestr, timestamp, url))
	return enpity.WxJsConfig{
		Appid:     appid,
		Timestmap: timestamp,
		NoceStr:   noncestr,
		Signature: signature,
	}
}

// 长链接转短链接接口，返回转换后的短链接结果
// 主要使用场景： 开发者用于生成二维码的原链接（商品、支付二维码等）太长导致扫码速度和成功率下降，
// 将原长链接通过此接口转成短链接再生成二维码将大大提升扫码速度和成功率。
func (w *WeChat) WxMakeShortLink(longUrl string) string {
	reqUrl := fmt.Sprintf(longurl_to_shorturl, w.WxGetAccessToken())
	body := map[string]string{
		"action": "long2short",
		"long_url": longUrl,
	}
	msg := http.Post(reqUrl, string(utils.Interface2byte(body)))
	var result map[string]string
	err := json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result["short_url"]
}

// 内部正式获取微信js的ticket信息
func getWxJsapiTicket(cfg *enpity.MpConfig) {
	reqUrl := fmt.Sprintf(jsapi_ticket, cfg.AccessToken)
	resp := http.Get(reqUrl)
	var jsTicket enpity.WxJsTicket
	err := json.Unmarshal(resp, &jsTicket)
	if err != nil {
		panic(err)
	}
	cfg.JsApiTicket = &jsTicket
}

// 检查微信功能调用的accessToken是否过期（注意，这里的accessToken不是获取用户信息的OAuthToken，而是调用微信api的token）
func isExpires(w WeChat) bool {
	return w.cfg.AccessTokenExpiresTime < time.Now().Unix()
}

// 内部调用刷新accessToken的微信api接口，此处是真正实现accessToken刷新的方法
func refreshToken(cfg *enpity.MpConfig, w WeChat) map[string]interface{} {
	requestUrl := fmt.Sprintf(access_token_url, cfg.AppId, cfg.Secret)
	msg := http.Get(requestUrl)
	var f interface{}
	json.Unmarshal(msg, &f)
	m := f.(map[string]interface{})
	return m
}