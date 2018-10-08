package error

import (
	"encoding/xml"
	"encoding/json"
	"fmt"
	"net/http"
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
	var tmp map[string]interface{}
	err := json.Unmarshal(result, &tmp)
	if err != nil {
		panic(err)
	}
	if val, ok := tmp["errcode"]; ok == true {
		if val != 0 || val != "0" {
			var wxErr WxMpError
			err = json.Unmarshal(result, &wxErr)
			if err != nil {
				panic(err)
			}
			panic(fmt.Errorf("微信调用api报错：[%s]", wxErr.ToJson(wxErr)))
		}
	}
}

func WxPayErrorFromByte(result []byte, resq *http.Response) []byte {
	return nil
}

