package service

import (
	"encoding/json"
	"fmt"
	"weixin-golang/weixin-common/http"
	"weixin-golang/weixin-mp/enpity"
)

const (
	set_industry_url			=		"https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s"
	get_industry_url			=		"https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=%s"
	add_template_url			=		"https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=%s"
	get_all_private_template	=		"https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%s"
	del_private_template		=		"https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=%s"
	send_template_msg			=		"https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"

)

// 设置所属行业
func (w *WeChat)WxTemplateSetIndustry(industry enpity.WxIndustry) string {
	reqUrl := fmt.Sprintf(set_industry_url, w.GetAccessToken())
	msg, err := http.Post(reqUrl, industry.ToJson(industry))
	if err != nil {
		panic(err)
	}
	return string(msg)
}

// 获取设置的行业信息
func (w *WeChat)WxTemplateGetIndustry() enpity.WxIndustryInfo {
	reqUrl := fmt.Sprintf(get_industry_url, w.GetAccessToken())
	msg, err := http.Get(reqUrl)
	if err != nil {
		panic(err)
	}
	var industryInfo enpity.WxIndustryInfo
	json.Unmarshal(msg, &industryInfo)
	return industryInfo
}


