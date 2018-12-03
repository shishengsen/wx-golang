package error

import (
	"encoding/json"
	"github.com/reactivex/rxgo/errors"
	"net/http"
)

type WxMpError struct {
	ErrCode				int32				`json:"errcode"`
	ErrMsg				string				`json:"errmsg"`
}

type WxPayError struct {
	ErrCode				int32				`json:"errcode"`
	ErrMsg				string				`json:"errmsg"`
}

func (e *WxMpError) Err() error {
	s, err := e.ToJson()
	if err != nil {
		return err
	}
	return errors.New(errors.ErrorCode(uint32(e.ErrCode)), s)
}

func (e *WxMpError)ToJson() (string, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// 对微信返回信息进行错误信息扫描，如果发现非正常状态返回，则抛出异常信息
func WxMpErrorFromByte(result []byte, resp *http.Response) (error) {
	var tmp WxMpError
	err := json.Unmarshal(result, &tmp)
	if err != nil {
		return err
	}
	if tmp.ErrCode != 0 {
		return tmp.Err()
	}
	return nil
}

func WxPayErrorFromByte(result []byte, resq *http.Response) []byte {
	return nil
}

