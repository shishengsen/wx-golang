package service

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/wxconsts"
	"wx-golang/weixin-mp/enpity"
	wxerr "wx-golang/weixin-common/error"
)

const (
	authorize_url           = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	oauth_access_token_url  = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	oauth_refresh_token_url = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	pull_user_info_url      = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s"
	verify_oauth_token      = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
)

// 根据 code 获取access_token
func (w *WeChat) getOAuthAccessToken(code string) enpity.WxOAuthAccessToken {
	reqUrl := fmt.Sprintf(oauth_access_token_url, w.Cfg.AppId, w.Cfg.Secret, code)
	msg := http.Get(reqUrl)
	var oauth enpity.WxOAuthAccessToken
	json.Unmarshal(msg, &oauth)
	oauth.ExpiresIn += +time.Now().Unix()
	oauth.IsExpires = false
	w.wxOAuthTokenStoreInMem(&oauth)
	return oauth
}

// 对外暴露的微信网页授权根据code获取accessToken
func (w *WeChat) GetOAuthTokenByCode(code string) enpity.WxOAuthAccessToken {
	return w.getOAuthAccessToken(code)
}

// 对外暴露的微信网页授权的accessToken
func (w *WeChat) GetOAuthToken() string {
	if w.Cfg.OAuthToken != nil {
		if !isOAuthExpires(){
			return w.Cfg.OAuthToken.OauthAccessToken
		}
	}
	return refreshOAuthToken().OauthAccessToken
}

// 对外暴露用户授权链接生成接口
func (w *WeChat) Oauth2buildAuthorizationUrl(redirectURI string, scope string, state string) string {
	redirectURI, err := url.QueryUnescape(redirectURI)
	if err != nil {
		panic(err)
	}
	authorizeUrl := fmt.Sprintf(authorize_url, w.Cfg.AppId, redirectURI, scope, state)
	return authorizeUrl
}

// 刷新token信息
func refreshOAuthToken() enpity.WxOAuthAccessToken {
	weChat := GetWeChat()
	if weChat.Cfg.OAuthToken.IsExpires {
		reqUrl := fmt.Sprintf(oauth_refresh_token_url, weChat.Cfg.AppId, weChat.Cfg.OAuthToken.OauthRefreshToken)
		msg := http.Get(reqUrl)
		var oauth enpity.WxOAuthAccessToken
		json.Unmarshal(msg, &oauth)
		oauth.ExpiresIn += +time.Now().Unix()
		oauth.IsExpires = false
		weChat.wxOAuthTokenStoreInMem(&oauth)
		return oauth
	}
	return *weChat.Cfg.OAuthToken
}

// 判断是否过期
func isOAuthExpires() bool {
	lock := sync.NewCond(new(sync.Mutex))
	lock.L.Lock()
	isExpires := GetWeChat().Cfg.OAuthToken.ExpiresIn-200 < time.Now().Unix()
	GetWeChat().Cfg.OAuthToken.IsExpires = isExpires
	lock.L.Unlock()
	return isExpires
}

// 检验授权凭证（access_token）是否有效
func (w *WeChat)VerifyToken(openid string) bool {
	reqUrl := fmt.Sprintf(verify_oauth_token, w.GetOAuthToken(), openid)
	result := http.Get(reqUrl)
	var returnCode wxerr.WxMpError
	wxerr.WxMpErrorFromByte(result, nil)
	json.Unmarshal(result, &returnCode)
	return returnCode.Errcode == 0
}

// 获取openid所对应的用户信息
func (w *WeChat)WxPullUserInfo(openid, lang string) enpity.WxMpUser {
	if lang == "" {
		lang = wxconsts.LANG_ZH_CN
	}
	reqUrl := fmt.Sprintf(pull_user_info_url, w.GetOAuthToken(), openid, lang)
	resp := http.Get(reqUrl)
	var opUser enpity.WxMpUser
	json.Unmarshal(resp, &opUser)
	return opUser
}

