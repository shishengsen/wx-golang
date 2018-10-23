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

// 对外暴露的微信网页授权根据code获取accessToken
func (w *WeChat) WxGetOAuthTokenByCode(code string) enpity.WxOAuthAccessToken {
	return getOAuthAccessToken(code, *w)
}

// 对外暴露的微信网页授权的accessToken
func (w *WeChat) WxGetOAuthToken() string {
	if w.cfg.OAuthToken != nil {
		if !isOAuthExpires(*w){
			return w.cfg.OAuthToken.OauthAccessToken
		}
	}
	return refreshOAuthToken(*w).OauthAccessToken
}

// 对外暴露用户授权链接生成接口
func (w *WeChat) WxOauth2buildAuthorizationUrl(redirectURI string, scope string, state string) string {
	return fmt.Sprintf(authorize_url, w.cfg.AppId, utils.UrlEncode(redirectURI), scope, state)
}


// 检验授权凭证（access_token）是否有效
func (w *WeChat) WxVerifyToken(openid string) bool {
	reqUrl := fmt.Sprintf(verify_oauth_token, w.WxGetOAuthToken(), openid)
	result := http.Get(reqUrl)
	var returnCode wxerr.WxMpError
	json.Unmarshal(result, &returnCode)
	return returnCode.Errcode == 0
}

// 获取openid所对应的用户信息
func (w *WeChat) WxPullUserInfo(openid, lang string) enpity.WxMpUser {
	if lang == "" {
		lang = wxconsts.LANG_ZH_CN
	}
	reqUrl := fmt.Sprintf(pull_user_info_url, w.WxGetOAuthToken(), openid, lang)
	resp := http.Get(reqUrl)
	var opUser enpity.WxMpUser
	json.Unmarshal(resp, &opUser)
	return opUser
}

// 根据 code 获取access_token
func getOAuthAccessToken(code string, w WeChat) enpity.WxOAuthAccessToken {
	reqUrl := fmt.Sprintf(oauth_access_token_url, w.cfg.AppId, w.cfg.Secret, code)
	msg := http.Get(reqUrl)
	var oauth enpity.WxOAuthAccessToken
	json.Unmarshal(msg, &oauth)
	oauth.ExpiresIn += +time.Now().Unix()
	oauth.IsExpires = false
	w.wxOAuthTokenStoreInMem(&oauth)
	return oauth
}

// 判断是否过期
func isOAuthExpires(wx WeChat) bool {
	lock := sync.NewCond(new(sync.Mutex))
	lock.L.Lock()
	isExpires := wx.cfg.OAuthToken.ExpiresIn-200 < time.Now().Unix()
	wx.cfg.OAuthToken.IsExpires = isExpires
	lock.L.Unlock()
	return isExpires
}

// 刷新token信息
func refreshOAuthToken(wx WeChat) enpity.WxOAuthAccessToken {
	if wx.cfg.OAuthToken.IsExpires {
		reqUrl := fmt.Sprintf(oauth_refresh_token_url, wx.cfg.AppId, wx.cfg.OAuthToken.OauthRefreshToken)
		msg := http.Get(reqUrl)
		var oauth enpity.WxOAuthAccessToken
		json.Unmarshal(msg, &oauth)
		oauth.ExpiresIn += +time.Now().Unix()
		oauth.IsExpires = false
		wx.wxOAuthTokenStoreInMem(&oauth)
		return oauth
	}
	return *wx.cfg.OAuthToken
}