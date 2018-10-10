package service

import (
	"encoding/json"
	"fmt"
	"wx-golang/weixin-common/http"
	wxerr "wx-golang/weixin-common/error"
	"wx-golang/weixin-mp/enpity"
)

const (
	/**
	创建标签
	 */
	mp_create_label			=			"https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s"
	/**
	获取公众号已创建的标签
	 */
	mp_label_list			=			"https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s"
)

// 创建标签
func (w *WeChat)WxMpUserCreateLabel(name string) string {
	reqUrl := fmt.Sprintf(mp_create_label, w.GetAccessToken())
	body := "{\"tag\":{\"name\":" + name + "}}"
	msg := http.Post(reqUrl, body)
	wxerr.WxMpErrorFromByte(msg, nil)
	var label map[string]map[string]string
	err := json.Unmarshal(msg, &label)
	if err != nil {
		panic(err)
	}
	return label["tag"]["id"]
}

// 获取公众号已创建的标签
func (w *WeChat)WxMpUserLabelList() enpity.WxMpUserLabel {
	reqUrl := fmt.Sprintf(mp_label_list, w.GetAccessToken())
	msg := http.Get(reqUrl)
	wxerr.WxMpErrorFromByte(msg, nil)
	var labels enpity.WxMpUserLabel
	err := json.Unmarshal(msg, &labels)
	if err != nil {
		panic(err)
	}
	return labels
}


