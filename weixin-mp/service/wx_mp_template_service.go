package service

import (
	"encoding/json"
	"fmt"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/utils"
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
func (w *WeChat) WxTemplateSetIndustry(industry enpity.WxIndustry) (string, error) {
	reqUrl := fmt.Sprintf(set_industry_url, w.WxGetAccessToken())
	msg, err := http.Post(reqUrl, industry.ToJson(industry))
	return string(msg), err
}

// 获取设置的行业信息
func (w *WeChat) WxTemplateGetIndustry() (enpity.WxIndustryInfo, error) {
	reqUrl := fmt.Sprintf(get_industry_url, w.WxGetAccessToken())
	msg, err := http.Get(reqUrl)
	var industryInfo enpity.WxIndustryInfo
	err = json.Unmarshal(msg, &industryInfo)
	return industryInfo, err
}

// 获得模板ID
func (w *WeChat) WxTemplateGetId(shortId string) (string, error) {
	reqUrl := fmt.Sprintf(add_template_url, w.WxGetAccessToken())
	body := map[string]string{
		"template_id_short": shortId,
	}
	msg, err := http.Post(reqUrl, string(utils.Interface2byte(body)))
	var responseBody map[string]string
	err = json.Unmarshal(msg, &responseBody)
	return responseBody["template_id"], err
}

// 获取模板列表
func (w *WeChat) WxTemplateGetTemplateList() (enpity.WxTemplateList, error) {
	reqUrl := fmt.Sprintf(get_all_private_template, w.WxGetAccessToken())
	msg, err := http.Get(reqUrl)
	var respBody enpity.WxTemplateList
	err = json.Unmarshal(msg, &respBody)
	return respBody, err
}

// 发送模板消息
func (w *WeChat) WxTemplateSendMsg(templateMsg enpity.WxTemplateMsg) (map[string]interface{}, error) {
	reqUrl := fmt.Sprintf(send_template_msg, w.WxGetAccessToken())
	msg, err := http.Post(reqUrl, string(utils.Interface2byte(templateMsg)))
	var respBody map[string]interface{}
	err = json.Unmarshal(msg, &respBody)
	return respBody, err
}
