package enpity

import (
	"encoding/json"
)

const (
	TYPE_CLICK				=			"click"
	TYPE_VIEW				=			"view"
	TYPE_SCANCODE_PUSH		=			"scancode_push"
	TYPE_SCANCODE_WAITMSG	=			"scancode_waitmsg"
	TYPE_PIC_SYSPHOTO		=			"pic_sysphoto"
	TYPE_PIC_PHOTO_OT_ALBUM =			"pic_photo_or_album"
	TYPE_PIC_WEIXIN			=			"pic_weixin"
	TYPE_LOCATION_SELECT	=			"location_select"
	TYPE_MEDIA_ID			=			"media_id"
	TYPE_VIEW_LIMITED		=			"view_limited"
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
func (this *WxMenu) ToJson() string {
	menu := WxMenu{}
	button := WxButton{
		Type:	TYPE_CLICK,
		Name:	"Test",
		Key:	"Key",
		Url:	"http://www.baidu.com",
	}
	sub := WxSubButton{
		Type:		TYPE_MEDIA_ID,
		MediaId:	"1",
	}
	button.Sub_button = []WxSubButton{sub}
	menu.Button = []WxButton{button}
	menuJson, err := json.Marshal(menu)
	if err != nil {
		panic(err)
	}
	return string(menuJson)
}
