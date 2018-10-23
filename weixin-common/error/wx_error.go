package error

import (
	"encoding/json"
	"net/http"
	"wx-golang/weixin-common/log"
)

type WxMpError struct {
	Errcode				int32				`json:"errcode"`
	Errmsg				string				`json:"errmsg"`
}

func (e *WxMpError)ToJson(w WxMpError) string {
	bytes, err := json.Marshal(w)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// 对微信返回信息进行错误信息扫描，如果发现非正常状态返回，则抛出异常信息
func WxMpErrorFromByte(result []byte, resp *http.Response) {
	var tmp WxMpError
	err := json.Unmarshal(result, &tmp)
	if err != nil {
		log.CheckError(err)
	}
	if tmp.Errcode != 0 {
		panic(tmp)
	}
}

func WxPayErrorFromByte(result []byte, resq *http.Response) []byte {

	return nil
}

