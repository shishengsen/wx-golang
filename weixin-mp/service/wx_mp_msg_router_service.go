package service

import (
	"wx-golang/weixin-mp/enpity"
)

// 微信消息通知路由分发
// 将自定义的微信路由处理函数装入map结构中
func (w *WeChat) RouterInit(msgRouters map[string]*enpity.MsgRouter) {
	w.handlerMap = msgRouters
}

// 路由分发，
func (w *WeChat) RouterSwitchHandler(msg enpity.WxMessage) {
	if &(w.handlerMap) == nil {
		panic("微信消息路由对象为空")
	}
	msgType := msg.MsgType
	router := w.handlerMap[msgType]
	(*router).Handler(msg)
}

