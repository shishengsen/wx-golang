package crypto

import (
	"sort"
	"encoding/hex"
	"crypto/sha1"
)

//sha1 签名
func Sha1(params ...string) string {
	sort.Strings(params)
	sha1 := sha1.New()
	var finalStr string
	for _, tmp := range params {
		finalStr += tmp
	}
	sha1.Write([]byte(finalStr))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

func WxCrypto() string {
	return ""
}