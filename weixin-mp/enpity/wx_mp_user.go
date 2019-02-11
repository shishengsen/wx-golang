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

type WxMpUserInfoBase struct {
	// 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	Subscribe int64 `json:"subscribe"`
	// 用户的标识，对当前公众号唯一
	Openid string `json:"openid"`
	// 用户的昵称
	Nickname string `json:"nickname"`
	// 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Sex int32 `json:"sex"`
	// 用户的语言，简体中文为zh_CN
	Language string `json:"language"`
	// 用户所在城市
	City string `json:"city"`
	// 用户所在省份
	Province string `json:"province"`
	// 用户所在国家
	Country string `json:"country"`
	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Headimgurl string `json:"headimgurl"`
	// 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	SubscribeTime int64 `json:"subscribe_time"`
	// 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Unionid string `json:"unionid"`
	// 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Remark string `json:"remark"`
	// 用户所在的分组ID（兼容旧的用户分组接口）
	Groupid int32 `json:"groupid"`
	// 用户被打上的标签ID列表
	TagidList []int32 `json:"tagid_list"`
	// 返回用户关注的渠道来源，
	// ADD_SCENE_SEARCH 公众号搜索，
	// ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，
	// ADD_SCENE_PROFILE_CARD 名片分享，
	// ADD_SCENE_QR_CODE 扫描二维码，
	// ADD_SCENEPROFILE LINK 图文页内名称点击，
	// ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，
	// ADD_SCENE_PAID 支付后关注，
	// ADD_SCENE_OTHERS 其他
	SubscribeScene string `json:"subscribe_scene"`
	// 二维码扫码场景（开发者自定义）
	QrScene int64 `json:"qr_scene"`
	// 二维码扫码场景描述（开发者自定义）
	QrSceneStr string `json:"qr_scene_str"`
}

type WxMpUserLabel struct {
	Tags []Label `json:"tags"`
}

type Label struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type OpenIdList struct {
	Total int64 `json:"total"`
	Count int64 `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

type WxMpUserList struct {
	Total      int64    `json:"total"`
	Count      int64    `json:"count"`
	Openid     []string `json:"openid"`
	NextOpenid []string `json:"next_openid"`
}
