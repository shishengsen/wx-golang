package example

import (
	"encoding/json"
	"fmt"
	"testing"
	"wx-golang/weixin-mp/enpity"
)

func TestKFJSON(t *testing.T)  {
	s := "{\"recordlist\":[{\"openid\":\"oDF3iY9WMaswOPWjCIp_f3Bnpljk\",\"opercode\":2002,\"text\":\" 您好，客服test1为您服务。\",\"time\":1400563710,\"worker\":\"test1@test\"},{\"openid\":\"oDF3iY9WMaswOPWjCIp_f3Bnpljk\",\"opercode\":2003,\"text\":\"你好，有什么事情？\",\"time\":1400563731,\"worker\":\"test1@test\"}],\"number\":2,\"msgid\":20165267}"
	var kf enpity.KfRecordList
	json.Unmarshal([]byte(s), &kf)
	k := enpity.KfRecordList{
		StartTime:987654321,
		EndTime:987654321,
		MsgId:1,
		Number:10000,
	}
	j, _ := json.Marshal(k)
	fmt.Printf("%#v\n", kf)
	fmt.Printf("%s", string(j))
}
