package service

import (
	"encoding/json"
	"fmt"
	"os"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/utils"
	"wx-golang/weixin-mp/enpity"
)

const (
	/**
	二维码请求
	*/
	qcode_request = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"
	/**
	通过ticket换取二维码
	*/
	ticket_to_qcode = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s"
)

// 请求临时二维码
func (w *WeChat) WxTempQCodeRequest(qrCodeReq enpity.WxQRCodeReq) (enpity.WxQRCodeResult, error) {
	body := map[string]interface{}{
		"expire_seconds": qrCodeReq.ExpireSeconds,
		"action_name":    qrCodeReq.ActionName,
		"action_info": map[string]interface{}{
			"scene": map[string]string{
				"scene_str": qrCodeReq.ActionInfo.SceneStr,
			},
		},
	}
	return qcodeRequest(*w, body)
}

// 请求永久二维码
func (w *WeChat) WxPermanentQCodeRequest(qrCodeReq enpity.WxQRCodeReq) (enpity.WxQRCodeResult, error) {
	body := map[string]interface{}{
		"action_name": qrCodeReq.ActionName,
		"action_info": map[string]interface{}{
			"scene": map[string]string{
				"scene_str": qrCodeReq.ActionInfo.SceneStr,
			},
		},
	}
	return qcodeRequest(*w, body)
}

// 通过ticket换取二维码
// 获取的是临时文件指针，需要自行将文件另存储
func (w *WeChat) WxTicket2QCode(ticket string) (*os.File, error) {
	reqUrl := fmt.Sprintf(ticket_to_qcode, utils.UrlEncode(ticket))
	msg, err := http.Get(reqUrl)
	return utils.CreateTempFile(msg), err
}

// 微信带参数二维码请求实际调用
func qcodeRequest(w WeChat, body interface{}) (enpity.WxQRCodeResult, error) {
	reqUrl := fmt.Sprintf(qcode_request, w.WxGetAccessToken())
	msg, err := http.Post(reqUrl, string(utils.Interface2byte(body)))
	var qCode enpity.WxQRCodeResult
	json.Unmarshal(msg, &qCode)
	return qCode, err
}
