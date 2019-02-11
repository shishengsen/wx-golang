package service

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	wxerr "wx-golang/weixin-common/error"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/utils"
	"wx-golang/weixin-common/wxconsts"
	"wx-golang/weixin-mp/enpity"
)

const (
	authorize_url           = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	oauth_access_token_url  = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	oauth_refresh_token_url = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	pull_user_info_url      = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s"
	verify_oauth_token      = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
)

type Token struct {
	wx		WeChat
	cfg		*enpity.MpConfig
}

// 刷新accessToken信息，token并发控制，防止过度刷新token信息导致access_token调用次数用完
func (t *Token) WxUpdateAccessToken() {
	t.cfg.AccessTokenLock.L.Lock()
	if isExpires(*t.cfg) {
		tokenMap, err := refreshToken(t.cfg)
		if err != nil {

		}
		t.cfg.AccessTokenExpiresTime = tokenMap["expires_in"].(int64) + time.Now().Unix()
		t.cfg.AccessToken = tokenMap["access_token"].(string)
	}
	t.cfg.AccessTokenLock.L.Unlock()
}

// 获取accessToken信息
func (t *Token) WxGetAccessToken() string {
	if isExpires(*t.cfg) {
		t.WxUpdateAccessToken()
	}
	return (*t.cfg).AccessToken
}

// 对外暴露的微信网页授权根据code获取accessToken
func (t *Token) WxGetOAuthTokenByCode(code string) enpity.WxOAuthAccessToken {
	return t.getOAuthAccessToken(code, *t.cfg)
}

// 对外暴露的微信网页授权的accessToken
func (t *Token) WxGetOAuthToken() string {
	if t.cfg.OAuthToken != nil {
		if !isOAuthExpires(t.cfg) {
			return t.cfg.OAuthToken.OauthAccessToken
		}
	}
	c, err := t.refreshOAuthToken(t.cfg)
	if err != nil {

	}
	return c.OauthAccessToken
}

// 对外暴露用户授权链接生成接口
func (t *Token) WxOauth2buildAuthorizationUrl(redirectURI string, scope string, state string) string {
	return fmt.Sprintf(authorize_url, t.cfg.AppId, utils.UrlEncode(redirectURI), scope, state)
}

// 检验授权凭证（access_token）是否有效
func (t *Token) WxVerifyOAuthToken(openid string) bool {
	reqUrl := fmt.Sprintf(verify_oauth_token, t.WxGetOAuthToken(), openid)
	result, err := http.Get(reqUrl)
	var returnCode wxerr.WxMpError
	json.Unmarshal(result, &returnCode)
	return err == nil
}

// 获取openid所对应的用户信息
func (t *Token) WxPullUserInfo(openid, lang string) enpity.WxMpUser {
	if lang == "" {
		lang = wxconsts.LANG_ZH_CN
	}
	reqUrl := fmt.Sprintf(pull_user_info_url, t.WxGetOAuthToken(), openid, lang)
	resp, err := http.Get(reqUrl)
	if err != nil {

	}
	var opUser enpity.WxMpUser
	json.Unmarshal(resp, &opUser)
	return opUser
}

// 根据 code 获取access_token
func (t *Token)getOAuthAccessToken(code string, cfg enpity.MpConfig) enpity.WxOAuthAccessToken {
	reqUrl := fmt.Sprintf(oauth_access_token_url, cfg.AppId, cfg.Secret, code)
	msg, err := http.Get(reqUrl)
	if err != nil {

	}
	var oauth enpity.WxOAuthAccessToken
	json.Unmarshal(msg, &oauth)
	oauth.ExpiresIn += +time.Now().Unix()
	oauth.IsExpires = false
	t.wx.wxOAuthTokenStoreInMem(&oauth)
	return oauth
}

// 判断是否过期
func isOAuthExpires(cfg *enpity.MpConfig) bool {
	lock := sync.NewCond(new(sync.Mutex))
	lock.L.Lock()
	isExpires := cfg.OAuthToken.ExpiresIn-200 < time.Now().Unix()
	cfg.OAuthToken.IsExpires = isExpires
	lock.L.Unlock()
	return isExpires
}

// 刷新token信息
func (t *Token)refreshOAuthToken(cfg *enpity.MpConfig) (enpity.WxOAuthAccessToken, error) {
	if cfg.OAuthToken.IsExpires {
		reqUrl := fmt.Sprintf(oauth_refresh_token_url, cfg.AppId, cfg.OAuthToken.OauthRefreshToken)
		msg, err := http.Get(reqUrl)
		var oauth enpity.WxOAuthAccessToken
		json.Unmarshal(msg, &oauth)
		oauth.ExpiresIn += +time.Now().Unix()
		oauth.IsExpires = false
		t.wx.wxOAuthTokenStoreInMem(&oauth)
		return oauth, err
	}
	return *cfg.OAuthToken, nil
}


// 检查微信功能调用的accessToken是否过期（注意，这里的accessToken不是获取用户信息的OAuthToken，而是调用微信api的token）
func isExpires(cfg enpity.MpConfig) bool {
	return cfg.AccessTokenExpiresTime < time.Now().Unix()
}

// 内部调用刷新accessToken的微信api接口，此处是真正实现accessToken刷新的方法
func refreshToken(cfg *enpity.MpConfig) (map[string]interface{}, error) {
	requestUrl := fmt.Sprintf(access_token_url, cfg.AppId, cfg.Secret)
	msg, err := http.Get(requestUrl)
	var f interface{}
	json.Unmarshal(msg, &f)
	m := f.(map[string]interface{})
	return m, err
}