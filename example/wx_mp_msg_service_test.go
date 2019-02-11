package example

import (
	"fmt"
	"testing"
	"wx-golang/weixin-common/wxconsts"
	"wx-golang/weixin-mp/enpity"
	"wx-golang/weixin-mp/service"
)


type GuanZhu struct{}

func (g *GuanZhu) Handler(enpity.WxMessage) {
	fmt.Print("关注事件")
}

type SaoMa struct{}

func (s *SaoMa) Handler(enpity.WxMessage) {
	fmt.Print("扫码事件")
}

func TestRouter(test *testing.T) {
	w := &service.WeChat{}
	router := w.RouterInit()
	g := GuanZhu{}
	router.Start().
		MsgType(wxconsts.MSG_TYPE_EVENT).
		Event(wxconsts.EVENT_TYPE_SUBSCRIBE).
		Handler(&g).
		End().
		Start().
		MsgType(wxconsts.MSG_TYPE_EVENT).
		Event(wxconsts.EVENT_TYPE_SCAN).
		Handler(&SaoMa{}).End()
	msg1 := enpity.WxMessage{MsgType:"event", Event:"subscribe"}
	msg2 := enpity.WxMessage{MsgType:"event", Event:"SCAN"}
	w.Route(msg1)
	w.Route(msg2)
}
