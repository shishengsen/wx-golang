package enpity

type WxQRCodeReq struct {
	ExpireSeconds string `json:"expire_seconds"`
	ActionName    string `json:"action_name"`
	ActionInfo    struct {
		SceneID  int32  `json:"scene_id"`
		SceneStr string `json:"scene_str"`
	} `json:"action_info"`
}

// 二维码ticket

type WxQRCodeResult struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	URL           string `json:"url"`
}
