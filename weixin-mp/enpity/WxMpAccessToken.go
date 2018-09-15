package enpity

type WxMpAccessToken struct {
	accessToken					string
	refreshToken				string
	accessTokenExpiresTime		int64
	refreshTokenExpiresTime		int64
}