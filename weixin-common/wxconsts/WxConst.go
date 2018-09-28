package wxconsts

const (

	// 微信网页授权

	SNSAPI_USERINFO = "snsapi_userinfo"
	SNSAPI_BASE     = "snsapi_base"

	// 微信菜单类别

	TYPE_CLICK              = "click"
	TYPE_VIEW               = "view"
	TYPE_SCANCODE_PUSH      = "scancode_push"
	TYPE_SCANCODE_WAITMSG   = "scancode_waitmsg"
	TYPE_PIC_SYSPHOTO       = "pic_sysphoto"
	TYPE_PIC_PHOTO_OT_ALBUM = "pic_photo_or_album"
	TYPE_PIC_WEIXIN         = "pic_weixin"
	TYPE_LOCATION_SELECT    = "location_select"
	TYPE_MEDIA_ID           = "media_id"
	TYPE_VIEW_LIMITED       = "view_limited"

	// 语言类别

	LANG_ZH_CN				= "zh_CN"
	LANG_ZH_TW				= "zh_TW"
	LANG_EN					= "en"

	// 微信网页授权链接

	AUTHORIZE_URL           = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	OAUTH_ACCESS_TOKEN_URL  = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	OAUTH_REFRESH_TOKEN_URL = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	PULL_USER_INFO_URL      = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	VERIFY_OAUTH_TOKEN      = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
)
