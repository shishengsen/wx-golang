package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 对微信返回信息进行错误信息扫描，如果发现非正常状态返回，则抛出异常信息
func WxMpError(result []byte, resp *http.Response) []byte {
	var tmp map[string]string
	err := json.Unmarshal(result, &tmp)
	if err != nil {
		panic(err)
	}
	if _, ok := tmp["errcode"]; ok == true {
		panic(fmt.Errorf("微信调用api报错：%#v", tmp))
	}
	return result
}

func WxPayError(result []byte, resq *http.Response) []byte {
	return nil
}
