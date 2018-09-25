package enpity

import "encoding/json"

// 设置所属行业
type WxIndustry struct {
	IndustryId1				string				`json:"industry_id1"`
	IndustryId2				string				`json:"industry_id2"`
}

// 获取设置的行业信息
type WxIndustryInfo struct {
	PrimaryIndustry				innerIndustry				`json:"primary_industry"`
	SecondaryIndustry			innerIndustry				`json:"secondary_industry"`
}

// 获取设置的行业信息——内部结构体
type innerIndustry struct {
	FirstClass				string				`json:"first_class"`
	SecondClass				string				`json:"second_class"`
}

// 获取模板列表
type WxTemplateList struct {
	TemplateList		[]WxTemplateInfo			`json:"template_list"`
}

// 模板信息
type WxTemplateInfo struct {
	TemplateId			string				`json:"template_id"`
	Title				string				`json:"title"`
	PrimaryIndustry		string				`json:"primary_industry"`
	DeputyIndustry		string				`json:"deputy_industry"`
	Content				string				`json:"content"`
	Example				string				`json:"example"`
}

// 模板消息
type WxTemplateMsg struct {
	ToUser				string				`json:"touser"`
	TemplateId			string				`json:"template_id"`
	Url					string				`json:"url"`
	Miniprogram			struct{
		AppId			string				`json:"appid"`
		Pagepath		string				`json:"pagepath"`
	}										`json:"miniprogram"`
	Data				struct{
		First			innerMsg			`json:"first"`
		Keyword1		innerMsg			`json:"keyword1"`
		Keyword2		innerMsg			`json:"keyword2"`
		Keyword3		innerMsg			`json:"keyword3"`
		Remark			innerMsg			`json:"remark"`
	}
}

// 消息设置
type innerMsg struct {
	Value				string				`json:"value"`
	Color				string				`json:"color"`
}

func (w *WxIndustry)ToJson(industry WxIndustry) string {
	industryJson, err := json.Marshal(industry)
	if err != nil {
		panic(err)
	}
	return string(industryJson)
}
