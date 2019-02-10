package service

import (
	"encoding/json"
	"fmt"
	"os"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-mp/enpity"
)

const (
	//
	add_kf_url = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s"
	//
	update_kf_url = "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s"
	//
	delete_kf_url = "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s"
	//
	set_kf_header = "http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%s&kf_account=%s"
	//
	get_kf_all = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s"
	//
	send_kf_msg = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
)

type WxKF struct {
	token	*Token
} 

// 添加客服信息
func (k *WxKF) AddKf(kf enpity.WxMpKf) (string, error) {
	reqUrl := fmt.Sprintf(add_kf_url, k.token.WxGetAccessToken())
	msg, err := http.Post(reqUrl, kf.ToJson(kf))
	return string(msg), err
}

// 更新客服信息
func (k *WxKF) UpdateKf(kf enpity.WxMpKf) (string, error) {
	reqUrl := fmt.Sprintf(update_kf_url, k.token.WxGetAccessToken())
	msg, err := http.Post(reqUrl, kf.ToJson(kf))
	return string(msg), err
}

// 删除客服信息
func (k *WxKF) DeleteKf(kf enpity.WxMpKf) (string, error) {
	reqUrl := fmt.Sprintf(delete_kf_url, k.token.WxGetAccessToken())
	msg, err := http.Post(reqUrl, kf.ToJson(kf))
	return string(msg), err
}

// 设置客服头像信息
func (k *WxKF) SetKfHeader(kf enpity.WxMpKf, file os.File) (string, error) {
	reqUrl := fmt.Sprintf(set_kf_header, k.token.WxGetAccessToken(), kf.KfAccount)
	msg, err := http.PostWithFile(reqUrl, &file)
	return string(msg), err
}

// 获取所有客服账号
func (k *WxKF) AllKf() (enpity.WxKfs, error) {
	reqUrl := fmt.Sprintf(get_kf_all, k.token.WxGetAccessToken())
	msg, err := http.Get(reqUrl)
	var kfs enpity.WxKfs
	json.Unmarshal(msg, &kfs)
	return kfs, err
}

// 发送客服消息
func (k *WxKF) SendKfMsg(kfMsg enpity.WxKfMsg) (string, error) {
	reqUrl := fmt.Sprintf(send_kf_msg, k.token.WxGetAccessToken())
	msg, err := http.Post(reqUrl, kfMsg.ToJson(kfMsg))
	return string(msg), err
}
