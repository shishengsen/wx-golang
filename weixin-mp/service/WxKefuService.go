package service

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"weixin-golang/weixin-common/http"
	"weixin-golang/weixin-mp/enpity"
)

const (
	add_kf_url			=			"https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s"
	update_kf_url		=			"https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s"
	delete_kf_url		=			"https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s"
	set_kf_header		=			"http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%s&kf_account=%s"
	get_kf_all			=			"https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s"
	send_kf_msg			=			"https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
)

// 添加客服信息
func AddKf(kf enpity.WxKf) string {
	req_url := fmt.Sprintf(add_kf_url, GetAccessToken())
	msg, err  := http.Post(req_url, kf.ToJson(kf))
	if err != nil {
		panic(err)
	}
	return string(msg)
}

// 更新客服信息
func UpdateKf(kf enpity.WxKf) string {
	req_url := fmt.Sprintf(update_kf_url, GetAccessToken())
	msg, err  := http.Post(req_url, kf.ToJson(kf))
	if err != nil {
		panic(err)
	}
	return string(msg)
}

// 删除客服信息
func DeleteKf(kf enpity.WxKf) string {
	req_url := fmt.Sprintf(delete_kf_url, GetAccessToken())
	msg, err := http.Post(req_url, kf.ToJson(kf))
	if err != nil {
		panic(err)
	}
	return string(msg)
}

// 设置客服头像信息
func SetKfHeader(kf enpity.WxKf, file multipart.File) string {
	req_url := fmt.Sprintf(set_kf_header, GetAccessToken(), kf.KfAccount)
	msg, err := http.PostWithFile(req_url, file)
	if err != nil {
		panic(err)
	}
	return string(msg)
}

// 获取所有客服账号
func AllKf() enpity.WxKfs {
	req_url := fmt.Sprintf(get_kf_all, GetAccessToken())
	msg, err := http.Get(req_url)
	if err != nil {
		panic(err)
	}
	var kfs enpity.WxKfs
	json.Unmarshal(msg, &kfs)
	return kfs
}

//
func SendKfMsg(kfMsg enpity.WxKfMsg) string {
	req_url := fmt.Sprintf(send_kf_msg, GetAccessToken())
	msg, err := http.Post(req_url, kfMsg.ToJson(kfMsg))
	if err != nil {
		panic(err)
	}
	return string(msg)
}