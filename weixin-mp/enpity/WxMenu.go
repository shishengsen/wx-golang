package enpity

import (
	"encoding/json"
)

type WxMenu struct {
	Button			[]WxButton		`json:"button,omitempty"`
}

type WxButton struct {
	Type			string			`json:"type,omitempty"`
	Name			string			`json:"name,omitempty"`
	Key				string			`json:"key,omitempty"`
	Url				string			`json:"url,omitempty"`
	MediaId			string			`json:"media_id,omitempty"`
	Sub_button		[]WxSubButton
}

type WxSubButton struct {
	Type			string			`json:"type,omitempty"`
	Name			string			`json:"name,omitempty"`
	Key				string			`json:"key,omitempty"`
	Url				string			`json:"url,omitempty"`
	MediaId			string			`json:"media_id,omitempty"`
}

// 将WxMenu转换为json字段
func (this *WxMenu) ToJson(menu WxMenu) string {
	menuJson, err := json.Marshal(menu)
	if err != nil {
		panic(err)
	}
	return string(menuJson)
}
