package enpity

import (
	"encoding/json"
)

// 微信自定义菜单
type WxMpMenu struct {
	Button    []WxButton `json:"button,omitempty"`
	Matchrule Matchrule  `json:"matchrule,omitempty"`
}

// 微信菜单主按钮
type WxButton struct {
	Type       string `json:"type,omitempty"`
	Name       string `json:"name,omitempty"`
	Key        string `json:"key,omitempty"`
	Url        string `json:"url,omitempty"`
	MediaId    string `json:"media_id,omitempty"`
	Sub_button []WxSubButton
}

// 微信菜单子按钮
type WxSubButton struct {
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
	Key     string `json:"key,omitempty"`
	Url     string `json:"url,omitempty"`
	MediaId string `json:"media_id,omitempty"`
}

// 微信个性化菜单的选项
type Matchrule struct {
	TagId              string `json:"tag_id,omitempty"`
	Sex                int32  `json:"sex,omitempty"`
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	ClientPlatformType string `json:"client_platform_type,omitempty"`
	Language           string `json:"language,omitempty"`
}

// 将WxMpMenu转换为json字段
func (this *WxMpMenu) ToJson(menu WxMpMenu) string {
	menuJson, err := json.Marshal(menu)
	if err != nil {
		panic(err)
	}
	return string(menuJson)
}
