package service

import (
	"time"
	"sync"
	"encoding/json"
	"fmt"
	"net/url"
	"weixin-golang/weixin-common/http"
	"weixin-golang/weixin-mp/enpity"
)

const (
	authorize_url			=	"https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	oauth_access_token_url	=	"https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	oauth_refresh_token_url	=	"https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	pull_user_info_url		=	"https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	verify_oauth_token		=	"https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
)

// 根据 code 获取access_token
func GetOAuthaccessToken(code string) enpity.WxOAuthAccessToken {
	cfg := GetMpConfig()
	req_url := fmt.Sprintf(oauth_access_token_url, cfg.AppId, cfg.Secret, code)
	msg, err := http.Get(req_url)
	if err != nil {
		panic(err)
	}
	var oauth enpity.WxOAuthAccessToken
	json.Unmarshal(msg, &oauth)
	oauth.ExpiresIn +=  + time.Now().Unix()
	oauth.IsExpires = false
	wxOAuthTokenStoreInMem(oauth)
	return oauth
}

// 对外暴露用户授权链接生成接口
func Oauth2buildAuthorizationUrl(redirectURI string, scope string, state string) string {
	redirectURI, err := url.QueryUnescape(redirectURI)
	if err != nil {
		panic(err)
	}
	cfg := GetMpConfig()
	authorizeUrl := fmt.Sprintf(authorize_url, cfg.AppId, redirectURI, scope, state)
	return authorizeUrl
}

// 刷新token信息
func RefreshOAuthToken() enpity.WxOAuthAccessToken {
	cfg := GetMpConfig()
	req_url := fmt.Sprintf(oauth_refresh_token_url, cfg.AppId, cfg.OAuthToken.OauthRefreshToken)
	msg, err := http.Get(req_url)
	if err != nil {
		panic(err)
	}
	var oauth enpity.WxOAuthAccessToken
	json.Unmarshal(msg, &oauth)
	oauth.ExpiresIn +=  + time.Now().Unix()
	oauth.IsExpires = false
	wxOAuthTokenStoreInMem(oauth)
	return oauth
}

// 判断是否过期
func isExpires() bool {
	lock := sync.NewCond(new(sync.Mutex))
	lock.L.Lock()
	isExpires := GetMpConfig().OAuthToken.ExpiresIn - 200 < time.Now().Unix()
	lock.L.Unlock()
	return isExpires
}

// 检验授权凭证（access_token）是否有效
func verifyToken(openid string) bool {
	cfg := GetMpConfig()
	req_url := fmt.Sprintf(verify_oauth_token, cfg.OAuthToken.OauthAccessToken, openid)
	_, err := http.Get(req_url)
	return err == nil
}
