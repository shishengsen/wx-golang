package service

import (
	"encoding/json"
	"fmt"
	"os"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/utils"
	wxerr "wx-golang/weixin-common/error"
	"wx-golang/weixin-mp/enpity"
)

const (
	/**
	二维码请求
	 */
	qcode_request 			=			"https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"
	/**
	通过ticket换取二维码
	 */
	ticket_to_qcode 		=			"https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s"
)

// 请求临时二维码
func (w *WeChat) WxTempQCodeRequest(expireSeconds int64, actionName, scene string) enpity.WxQCode {
	body := map[string]interface{} {
		"expire_seconds": expireSeconds,
		"action_name": actionName,
		"action_info": map[string]interface{} {
			"scene": map[string]string {
				"scene_str": scene,
			},
		},
	}
	return qcodeRequest(*w, body)
}

// 请求永久二维码
func (w *WeChat) WxPermanentQCodeRequest(actionName, scene string) enpity.WxQCode {
	body := map[string]interface{} {
		"action_name": actionName,
		"action_info": map[string]interface{} {
			"scene": map[string]string {
				"scene_str": scene,
			},
		},
	}
	return qcodeRequest(*w, body)
}

// 通过ticket换取二维码
// 获取的是临时文件指针，需要自行将文件另存储
func (w *WeChat) WxTicket2QCode(ticket string) *os.File {
	reqUrl := fmt.Sprintf(ticket_to_qcode, utils.UrlEncode(ticket))
	msg := http.Get(reqUrl)
	return utils.CreateTempFile(msg)
}

// 微信带参数二维码请求实际调用
func qcodeRequest(w WeChat, body interface{}) enpity.WxQCode {
	reqUrl := fmt.Sprintf(qcode_request, w.WxGetAccessToken())
	msg := http.Post(reqUrl, string(utils.Interface2byte(body)))
	var qCode enpity.WxQCode
	err := json.Unmarshal(msg, &qCode)
	if err != nil {
		panic(err)
	}
	return qCode
}
