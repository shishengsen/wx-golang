package service

import (
	"weixin-golang/weixin-mp/enpity"
)

var handlerMap map[string]*enpity.MsgRouter

// 微信消息通知路由分发
func (w *WeChat) RouterInit(msgRouters map[string]*enpity.MsgRouter) {
	handlerMap = msgRouters
}

// 路由分发
func (w *WeChat) RouterSwitchHandler(msg enpity.WxMessage) {
	msgType := msg.MsgType
	router := handlerMap[msgType]
	(*router).Handler(msg)
}

