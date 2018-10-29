package enpity

import "time"

type WxJsConfig struct {
	Appid			string				`json:"appid"`
	Timestmap		int64				`json:"timestmap"`
	NoceStr			string				`json:"noceStr"`
	Signature		string				`json:"signature"`
}

type WxJsTicket struct {
	Ticket			string				`json:"ticket"`
	ExpiresIn		int64				`json:"expires_in"`
}

func (js *WxJsTicket) IsExpires() bool {
	return js.ExpiresIn < time.Now().Unix()
}