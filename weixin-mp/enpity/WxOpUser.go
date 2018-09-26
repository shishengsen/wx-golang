package enpity

type WxOpUser struct {
	OpenId     string   `json:"openid"`
	NickName   string   `json:"nickname"`
	Sex        int32    `json:"sex"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}
