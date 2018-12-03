package utils

import (
	"github.com/panjf2000/ants"
	"wx-golang/weixin-mp/enpity"
)

var mpMsgHandlerPool *ants.Pool
var err error

func init() {
	mpMsgHandlerPool, err = ants.NewPool(10)
	if err != nil {
		panic(err)
	}
}

func SubmitMpMsgWork(wxMsg enpity.WxMessage, fun func(enpity.WxMessage)) {
	_ = mpMsgHandlerPool.Submit(func() error {
		fun(wxMsg)
		return nil
	})
}
