package enpity

type WxMpUser struct {
	OpenId     string   `json:"openid"`
	NickName   string   `json:"nickname"`
	Sex        int32    `json:"sex"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

type WxMpUserLabel struct {
	Tags		[]Label		`json:"tags"`
}

type Label struct {
	Id			string		`json:"id"`
	Name		string		`json:"name"`
}

type LabelFans struct {
	Count		int64			`json:"count"`
	Data		struct{
		Openid		[]string	`json:"openid"`
	}							`json:"data"`
	NextOpenid	string			`json:"next_openid"`
}
