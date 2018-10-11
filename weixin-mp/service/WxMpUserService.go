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
	mp_create_label				=			"https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s"
	/**
	获取公众号已创建的标签
	 */
	mp_label_list				=			"https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s"
	/**
	编辑标签
	 */
	mp_label_update				=			"https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s"
	/**
	删除标签
	 */
	mp_label_delete				=			"https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s"
	/**
	获取标签下粉丝列表
	 */
	label_fans_list				=			"https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s"
	/**
	批量为用户打标签
	 */
	mp_batch_user_label_add 	=			"https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=%s"
	/**
	批量为用户取消标签
	 */
	mp_batch_user_label_cancle	=			"https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=%s"
	/**
	获取用户身上的标签列表
	 */
	mp_user_own_labels			=			"https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=%s"
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

// 编辑标签
func (w *WeChat)WxMpUserLabelUpdate(label enpity.Label) bool {
	reqUrl := fmt.Sprintf(mp_label_update, w.GetAccessToken())
	body, err := json.Marshal(label)
	if err != nil {
		panic(err)
	}
	msg := http.Post(reqUrl, string(body))
	wxerr.WxMpErrorFromByte(msg, nil)
	return true
}

// 删除标签
func (w *WeChat)WxMpUserLabelDelete(id int32) bool {
	reqUrl := fmt.Sprintf(mp_label_delete, w.GetAccessToken())
	body := "{\"tag\":{\"id\": " + string(id) + "}}"
	msg := http.Post(reqUrl, body)
	wxerr.WxMpErrorFromByte(msg, nil)
	return true
}

// 获取标签下粉丝列表
func (w *WeChat)WxMpLabelFansList() enpity.LabelFans {
	reqUrl := fmt.Sprintf(label_fans_list, w.GetAccessToken())
	msg := http.Get(reqUrl)
	wxerr.WxMpErrorFromByte(msg, nil)
	var fans enpity.LabelFans
	err := json.Unmarshal(msg, &fans)
	if err != nil {
		panic(err)
	}
	return fans
}

// 批量为用户打标签
func (w *WeChat) WxMpUserBatchLabelAdd(openIds []string, tagId int32) bool {
	reqUrl := fmt.Sprintf(mp_batch_user_label_add, w.GetAccessToken())
	return mpUserBatchLabel(reqUrl, openIds, tagId)
}

// 批量为用户取消标签
func (w *WeChat)WxMpUserBatchLabelCancle(openIds []string, tagId int32) bool {
	reqUrl := fmt.Sprintf(mp_batch_user_label_cancle, w.GetAccessToken())
	return mpUserBatchLabel(reqUrl, openIds, tagId)
}

// 获取用户身上的标签列表
func (w *WeChat)WxMpUserOwnLabels(openId string) []string {
	reqUrl := fmt.Sprintf(mp_user_own_labels, w.GetAccessToken())
	body := "{\"openid\":" + openId + "}"
	msg := http.Post(reqUrl, body)
	wxerr.WxMpErrorFromByte(msg, nil)
	var result map[string][]string
	err := json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result["tagid_list"]
}

func mpUserBatchLabel(reqUrl string, openIds []string, tagId int32) bool {
	body := map[string]interface{}{
		"openid_list": openIds,
		"tagid": tagId,
	}
	bodyByte, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	msg := http.Post(reqUrl, string(bodyByte))
	wxerr.WxMpErrorFromByte(msg, nil)
	return true
}


