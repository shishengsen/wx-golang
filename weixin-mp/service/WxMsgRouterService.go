package service

import (
	"weixin-golang/weixin-mp/enpity"
)

const (
	MSG_TYPR_TEXT			=		"text"
	MSG_TYPE_IMAGE			=		"image"
	MSG_TYPE_VOICE			=		"voice"
	MSG_TYPE_VIDEO			=		"video"
	MSG_TYPE_SHORT_VIDEO	=		"shortvideo"
	MSG_TYPE_LOCATION		=		"location"
	MSG_TYPE_LINK			=		"link"
	MSG_TYPE_EVENT			=		"event"
)

var handlerMap map[string]*enpity.MsgRouter

// 微信消息通知路由分发
func (w *WeChat) RouterInit(msgRouters map[string]*enpity.MsgRouter) {
	handlerMap = msgRouters
}

// 路由分发，
func (w *WeChat) RouterSwitchHandler(msg enpity.WxMessage) {
	if &handlerMap == nil {
		panic("微信消息路由对象为空")
	}
	msgType := msg.MsgType
	router := handlerMap[msgType]
	(*router).Handler(msg)
}

