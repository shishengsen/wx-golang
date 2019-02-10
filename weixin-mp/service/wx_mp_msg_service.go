package service

import (
	"wx-golang/weixin-common/utils"
	"wx-golang/weixin-mp/enpity"
)

// 继承该接口实现微信消息的处理
type MsgHandler interface {
	Handler(enpity.WxMessage)
}

// 微信消息路由结构体
type MsgRouter struct {
	rules []*MsgRule
}

// 路由消息匹配规则
type MsgRule struct {
	router   *MsgRouter
	fromUser string
	msgType  string
	event    string
	eventKey string
	content  string
	async    bool
	handler  MsgHandler
}

// 初始化消息路由
func (w *WeChat) RouterInit() *MsgRouter {
	return &MsgRouter{}
}

// 开始路由规则匹配
func (r *MsgRouter) Start() *MsgRule {
	return &MsgRule{router: r}
}

// 结束路由规则匹配
func (r *MsgRule) End() *MsgRouter {
	r.router.rules = append(r.router.rules, r)
	return r.router
}

// 微信消息路由分发消息处理
// 目前设置每个路由消息处理仅处理一次
// 根据Async决定是否使用异步操作
func (w *WeChat) Route(wxMsg enpity.WxMessage) {
	routers := w.Router
	for _, rule := range routers.rules {
		if rule.match(wxMsg) {
			if rule.async {
				utils.SubmitMpMsgWork(wxMsg, rule.handler.Handler)
			} else {
				rule.handler.Handler(wxMsg)
			}
		}
	}
}

// ###################### 消息匹配规则 ######################

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

func (r *MsgRule) Handler(handler MsgHandler) *MsgRule {
	r.handler = handler
	return r
}

// 某个消息事件多个匹配规则（待支持）
func (r *MsgRule) Next() *MsgRule {
	return r
}

// 消息规则匹配
func (r *MsgRule) match(wxMsg enpity.WxMessage) bool {
	return r.msgType == wxMsg.MsgType &&
		r.event == wxMsg.Event &&
		r.eventKey == wxMsg.EventKey &&
		r.content == wxMsg.Content
}
