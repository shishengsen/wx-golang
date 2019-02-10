package service

import (
	"fmt"
	"testing"
	"wx-golang/weixin-mp/enpity"
)

type GuanZhu struct{}

func (g *GuanZhu) Handler(enpity.WxMessage) {
	fmt.Print("关注事件")
}

type SaoMa struct{}

func (s *SaoMa) Handler(enpity.WxMessage) {
	fmt.Print("扫码事件")
}

func TeestRouter(test *testing.T) {
	router := &MsgRouter{}
	router.Begin()
}
