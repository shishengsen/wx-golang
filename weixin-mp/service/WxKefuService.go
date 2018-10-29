package service

import (
	"encoding/json"
	"fmt"
	"os"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-mp/enpity"
)

const (
	add_kf_url    = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s"
	update_kf_url = "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s"
	delete_kf_url = "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s"
	set_kf_header = "http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%s&kf_account=%s"
	get_kf_all    = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s"
	send_kf_msg   = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
)

// 添加客服信息
func (w *WeChat) AddKf(kf enpity.WxKf) string {
	reqUrl := fmt.Sprintf(add_kf_url, w.GetAccessToken())
	msg := http.Post(reqUrl, kf.ToJson(kf))
	return string(msg)
}

// 更新客服信息
func (w *WeChat) UpdateKf(kf enpity.WxKf) string {
	reqUrl := fmt.Sprintf(update_kf_url, w.GetAccessToken())
	msg := http.Post(reqUrl, kf.ToJson(kf))
	return string(msg)
}

// 删除客服信息
func (w *WeChat) DeleteKf(kf enpity.WxKf) string {
	reqUrl := fmt.Sprintf(delete_kf_url, w.GetAccessToken())
	msg := http.Post(reqUrl, kf.ToJson(kf))
	return string(msg)
}

// 设置客服头像信息
func (w *WeChat) SetKfHeader(kf enpity.WxKf, file os.File) string {
	reqUrl := fmt.Sprintf(set_kf_header, w.GetAccessToken(), kf.KfAccount)
	msg := http.PostWithFile(reqUrl, &file)
	return string(msg)
}

// 获取所有客服账号
func (w *WeChat) AllKf() enpity.WxKfs {
	reqUrl := fmt.Sprintf(get_kf_all, w.GetAccessToken())
	msg := http.Get(reqUrl)
	var kfs enpity.WxKfs
	json.Unmarshal(msg, &kfs)
	return kfs
}

// 发送客服消息
func (w *WeChat) SendKfMsg(kfMsg enpity.WxKfMsg) string {
	reqUrl := fmt.Sprintf(send_kf_msg, w.GetAccessToken())
	msg := http.Post(reqUrl, kfMsg.ToJson(kfMsg))
	return string(msg)
}