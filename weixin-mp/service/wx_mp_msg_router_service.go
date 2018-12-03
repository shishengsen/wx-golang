package service

import (
	"ProjectOne/routers"
	"wx-golang/weixin-mp/enpity"
)

// 继承该接口即可实现微信信息路由分发
type MsgHandler interface {
	Handler(WxMessage)
}

// 路由消息匹配规则
type MsgRule struct {
	router				*MsgRouter
	fromUser			string
	msgType				string
	event				string
	eventKey			string
	content				string
	async				bool
	handler				MsgHandler
}

// 微信消息路由结构体
type MsgRouter struct {
	rules			[]MsgRule
}

// 初始化消息路由
func (w *WeChat) RouterInit() *MsgRouter {
	return &MsgRouter{}
}

// 开始路由规则匹配
func (r *MsgRouter) Begin(route *MsgRouter) *MsgRule {
	return &MsgRule{router: route}
}

// 结束路由规则匹配
func (r *MsgRule) End() *MsgRouter {
	r.router.rules = append(r.router.rules, r)
}

func (r *MsgRule) FromUser(fromUser string) *MsgRule {
	r.fromUser = fromUser
	return r
}

func (r *MsgRule) MsgType(msgType string) *MsgRule {
	r.msgType = msgType
	return r
}

func (r *MsgRule) Event(event string) *MsgRule {
	r.event = event
	return r
}

func (r *MsgRule) EventKey(eventKey string) *MsgRule {
	r.eventKey = eventKey
	return r
}

func (r *MsgRule) Content(content string) *MsgRule {
	r.content = content
	return r
}

func (r *MsgRule) Async(async bool) *MsgRule {
	r.async = async
	return r
}

func (r *MsgRule) match(wxMsg enpity.WxMessage) bool {
	return r.fromUser == wxMsg.FromUserName && r.msgType == wxMsg.MsgType && r.event == wxMsg.Event && r.eventKey == wxMsg.EventKey
}


