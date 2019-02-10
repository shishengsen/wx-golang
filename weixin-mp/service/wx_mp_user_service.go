package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/utils"
	"wx-golang/weixin-mp/enpity"
)

const (
	/**
	创建标签
	*/
	mp_create_label = "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s"
	/**
	获取公众号已创建的标签
	*/
	mp_label_list = "https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s"
	/**
	编辑标签
	*/
	mp_label_update = "https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s"
	/**
	删除标签
	*/
	mp_label_delete = "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s"
	/**
	获取标签下粉丝列表
	*/
	label_fans_list = "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s"
	/**
	批量为用户打标签
	*/
	mp_batch_user_label_add = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=%s"
	/**
	批量为用户取消标签
	*/
	mp_batch_user_label_cancle = "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=%s"
	/**
	获取用户身上的标签列表
	*/
	mp_user_own_labels = "https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=%s"
	/**
	获取用户基本信息（包括UnionID机制）
	*/
	mp_user_info_base = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=%s"
	/**
	批量获取用户基本信息
	*/
	mp_user_info_base_list = "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=%s"
	/**
	获取用户列表
	*/
	mp_user_list = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s"
	/**
	获取公众号的黑名单列表
	*/
	mp_user_black_list = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=%s"
)

// 创建标签
func (w *WeChat) WxMpUserCreateLabel(name string) string {
	reqUrl := fmt.Sprintf(mp_create_label, w.WxGetAccessToken())
	body := "{\"tag\":{\"name\":" + name + "}}"
	msg, err := http.Post(reqUrl, body)
	var label map[string]map[string]string
	err = json.Unmarshal(msg, &label)
	if err != nil {
		panic(err)
	}
	return label["tag"]["id"]
}

// 获取公众号已创建的标签
func (w *WeChat) WxMpUserLabelList() enpity.WxMpUserLabel {
	reqUrl := fmt.Sprintf(mp_label_list, w.WxGetAccessToken())
	msg, err := http.Get(reqUrl)
	var labels enpity.WxMpUserLabel
	err = json.Unmarshal(msg, &labels)
	if err != nil {
		panic(err)
	}
	return labels
}

// 编辑标签
func (w *WeChat) WxMpUserLabelUpdate(label enpity.Label) bool {
	reqUrl := fmt.Sprintf(mp_label_update, w.WxGetAccessToken())
	_, err := http.Post(reqUrl, string(utils.Interface2byte(label)))
	if err != nil {
	}
	return true
}

// 删除标签
func (w *WeChat) WxMpUserLabelDelete(id int32) bool {
	reqUrl := fmt.Sprintf(mp_label_delete, w.WxGetAccessToken())
	body := "{\"tag\":{\"id\": " + string(id) + "}}"
	_, err := http.Post(reqUrl, body)
	if err != nil {

	}
	return true
}

// 获取标签下粉丝列表
func (w *WeChat) WxMpLabelFansList() enpity.OpenIdList {
	reqUrl := fmt.Sprintf(label_fans_list, w.WxGetAccessToken())
	msg, err := http.Get(reqUrl)
	var fans enpity.OpenIdList
	err = json.Unmarshal(msg, &fans)
	if err != nil {
		panic(err)
	}
	return fans
}

// 批量为用户打标签
func (w *WeChat) WxMpUserBatchLabelAdd(openIds []string, tagId int32) (bool, error) {
	reqUrl := fmt.Sprintf(mp_batch_user_label_add, w.WxGetAccessToken())
	return mpUserBatchLabel(reqUrl, openIds, tagId)
}

// 批量为用户取消标签
func (w *WeChat) WxMpUserBatchLabelCancle(openIds []string, tagId int32) (bool, error) {
	reqUrl := fmt.Sprintf(mp_batch_user_label_cancle, w.WxGetAccessToken())
	return mpUserBatchLabel(reqUrl, openIds, tagId)
}

// 获取用户身上的标签列表
func (w *WeChat) WxMpUserOwnLabels(openId string) []string {
	reqUrl := fmt.Sprintf(mp_user_own_labels, w.WxGetAccessToken())
	body := "{\"openid\":" + openId + "}"
	msg, err := http.Post(reqUrl, body)
	var result map[string][]string
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result["tagid_list"]
}

// 获取用户基本信息（包括UnionID机制）
func (w *WeChat) WxMpGetUserInfoByUUID(openId string) enpity.WxMpUserInfoBase {
	reqUrl := fmt.Sprintf(mp_user_info_base, w.WxGetAccessToken(), openId, w.cfg.Lang)
	msg, err := http.Get(reqUrl)
	var userBaseInfo enpity.WxMpUserInfoBase
	err = json.Unmarshal(msg, &userBaseInfo)
	if err != nil {
		panic(err)
	}
	return userBaseInfo
}

// 批量获取用户基本信息
func (w *WeChat) WxMpGetUserInfoList(openIds []string) []enpity.WxMpUserInfoBase {
	reqUrl := fmt.Sprintf(mp_user_info_base_list, w.WxGetAccessToken())
	userList := make([]map[string]string, len(openIds))
	for i := range openIds {
		_tmp := map[string]string{
			"openid": openIds[i],
			"lang":   w.cfg.Lang,
		}
		userList = append(userList, _tmp)
	}
	msg, err := http.Post(reqUrl, string(utils.Interface2byte(userList)))
	var userInfoBases []enpity.WxMpUserInfoBase
	err = json.Unmarshal(msg, &userInfoBases)
	if err != nil {
		panic(err)
	}
	return userInfoBases
}

// 获取用户列表
func (w *WeChat) WxMpGetUserList(nextOpenId string) enpity.WxMpUserList {
	reqUrl := fmt.Sprintf(mp_user_list, w.WxGetAccessToken())
	if strings.Compare(nextOpenId, "") != 0 {
		reqUrl += "&next_openid=" + nextOpenId
	}
	msg, err := http.Get(reqUrl)
	var userList enpity.WxMpUserList
	err = json.Unmarshal(msg, &userList)
	if err != nil {
		panic(err)
	}
	return userList
}

// 获取公众号的黑名单列表
func (w *WeChat) WxMpGetUserBlacklist(openId string) enpity.OpenIdList {
	reqUrl := fmt.Sprintf(mp_user_black_list, w.WxGetAccessToken())
	msg, err := http.Post(reqUrl, "{\"begin_openid\":"+openId+"}")
	var blacklist enpity.OpenIdList
	err = json.Unmarshal(msg, &blacklist)
	if err != nil {
		panic(err)
	}
	return blacklist
}

//
func mpUserBatchLabel(reqUrl string, openIds []string, tagId int32) (bool, error) {
	body := map[string]interface{}{
		"openid_list": openIds,
		"tagid":       tagId,
	}
	bodyByte, _ := json.Marshal(body)
	_, err := http.Post(reqUrl, string(bodyByte))
	if err != nil {
		return false, err
	}
	return true, nil
}
