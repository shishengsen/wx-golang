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
	access_token_url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	/**

	 */
	clear_quota = "https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=%s"
	/**

	 */
	jsapi_ticket = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	/**

	 */
	jsapi_signature = "jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s"
	/**
	长链接转短链接接口
	*/
	longurl_to_shorturl = "https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s"
)

type WeChat struct {
	cfg    	*enpity.MpConfig
	Router 	*MsgRouter
	Token	*Token
	Menu	*WxMenu
	Data	*WxDataAnalyze
	User	*WxUser
	KF		*WxKF
	Material		*WxMaterial
}

// 确保只初始化一次 MpConfig
func WxNewMpConfig(cfg *enpity.MpConfig) *WeChat {
	cfg.IsExpire = false
	err := enpity.Validator(cfg)
	if err != nil {
		panic(err)
	}
	weChat := new(WeChat)
	cfg.AccessTokenLock = sync.NewCond(new(sync.Mutex))
	cfg.JsapiTicketLock = sync.NewCond(new(sync.Mutex))
	weChat.cfg = cfg
	return weChat
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

// 公众号调用或第三方平台帮公众号调用对公众号的所有api调用（包括第三方帮其调用）次数进行清零：
func (w *WeChat) WxApiClearQuota() (map[string]interface{}, error) {
	reqUrl := fmt.Sprintf(clear_quota, w.Token.WxGetAccessToken())
	reqBody := map[string]string{
		"appid": w.cfg.AppId,
	}
	msg, err := http.Post(reqUrl, string(utils.Interface2byte(reqBody)))
	var respBody map[string]interface{}
	json.Unmarshal(msg, &reqBody)
	return respBody, err
}

// 微信通知事件分发
func (w *WeChat) WxMpMsgPushEvent(s string) {
	var msg enpity.WxMessage
	_ = xml.Unmarshal([]byte(s), &msg)
	w.Route(msg)
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
	reqUrl := fmt.Sprintf(longurl_to_shorturl, w.Token.WxGetAccessToken())
	body := map[string]string{
		"action":   "long2short",
		"long_url": longUrl,
	}
	msg, err := http.Post(reqUrl, string(utils.Interface2byte(body)))
	var result map[string]string
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result["short_url"]
}

// 内部正式获取微信js的ticket信息
func getWxJsapiTicket(cfg *enpity.MpConfig) {
	reqUrl := fmt.Sprintf(jsapi_ticket, cfg.AccessToken)
	resp, err := http.Get(reqUrl)
	var jsTicket enpity.WxJsTicket
	err = json.Unmarshal(resp, &jsTicket)
	if err != nil {
		panic(err)
	}
	cfg.JsApiTicket = &jsTicket
}
