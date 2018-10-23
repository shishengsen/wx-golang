package enpity

type WxOAuthAccessToken struct {
	OauthAccessToken  string `json:"access_token"`
	OauthRefreshToken string `json:"refresh_token"`
	ExpiresIn         int64  `json:"expires_in"`
	OpenId            string `json:"openid"`
	Scope             string `json:"scope"`
	IsExpires         bool
}
