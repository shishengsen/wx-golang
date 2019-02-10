package enpity

// 二维码请求

type WxQRCodeReq struct {
	ExpireSeconds string `json:"expire_seconds"`
	ActionName    string `json:"action_name"`
	ActionInfo    struct {
		SceneID  int32  `json:"scene_id"`
		SceneStr string `json:"scene_str"`
	} `json:"action_info"`
}

// 二维码获取结果

type WxQRCodeResult struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	URL           string `json:"url"`
}
