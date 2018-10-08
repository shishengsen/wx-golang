package service

import (
	"encoding/json"
	"fmt"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-mp/enpity"
)

const (
	set_industry_url         = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s"
	get_industry_url         = "https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=%s"
	add_template_url         = "https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=%s"
	get_all_private_template = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%s"
	del_private_template     = "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=%s"
	send_template_msg        = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
)

// 设置所属行业
func (w *WeChat) WxTemplateSetIndustry(industry enpity.WxIndustry) string {
	reqUrl := fmt.Sprintf(set_industry_url, w.GetAccessToken())
	msg := http.Post(reqUrl, industry.ToJson(industry))
	return string(msg)
}

// 获取设置的行业信息
func (w *WeChat) WxTemplateGetIndustry() enpity.WxIndustryInfo {
	reqUrl := fmt.Sprintf(get_industry_url, w.GetAccessToken())
	msg := http.Get(reqUrl)
	var industryInfo enpity.WxIndustryInfo
	json.Unmarshal(msg, &industryInfo)
	return industryInfo
}

// 获得模板ID
func (w *WeChat) WxTemplateGetId(shortId string) string {
	reqUrl := fmt.Sprintf(add_template_url, w.GetAccessToken())
	body, err := json.Marshal(map[string]string{
		"template_id_short": shortId,
	})
	if err != nil {
		panic(err)
	}
	msg := http.Post(reqUrl, string(body))
	var responseBody map[string]string
	json.Unmarshal(msg, &responseBody)
	return responseBody["template_id"]
}

// 获取模板列表
func (w *WeChat) WxTemplateGetTemplateList() enpity.WxTemplateList {
	reqUrl := fmt.Sprintf(get_all_private_template, w.GetAccessToken())
	msg := http.Get(reqUrl)
	var respBody enpity.WxTemplateList
	json.Unmarshal(msg, &respBody)
	return respBody
}

// 发送模板消息
func (w *WeChat) WxTemplateSendMsg(templateMsg enpity.WxTemplateMsg) map[string]interface{} {
	reqUrl := fmt.Sprintf(send_template_msg, w.GetAccessToken())
	reqBody, err := json.Marshal(templateMsg)
	if err != nil {
		panic(err)
	}
	msg := http.Post(reqUrl, string(reqBody))
	var respBody map[string]interface{}
	json.Unmarshal(msg, &respBody)
	return respBody
}
