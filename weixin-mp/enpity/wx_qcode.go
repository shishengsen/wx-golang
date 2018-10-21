package enpity

// 二维码ticket
type WxQCode struct {
	Ticket			string				`json:"ticket"`
	ExpireSeconds	int64				`json:"expire_seconds"`
	Url				string				`json:"url"`
}