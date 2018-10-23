package crypto

// 此处AES代码参考——(golang使用aes库实现加解密)[https://blog.csdn.net/robertkun/article/details/79218088]

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"socket/util"
	"sort"
)

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// AES消息加密
func aesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = pKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AES消息解密
func aesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS7UnPadding(origData)
	return origData, nil
}

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

//sha1签名(已拼接好的字符串)
func Sha1WithAmple(str string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

// 微信消息加密
func WxMsgAesEncrypt(cipherText string, aesKey string) string {
	cipherByte := []byte(cipherText)
	aesKeyByte := []byte(aesKey)
	xpass, err := aesEncrypt(cipherByte, aesKeyByte)
	if err != nil {
		panic(err)
	}

	pass64 := base64.StdEncoding.EncodeToString(xpass)
	util.GetLogger().Infof("加密后:%v\n", pass64)
	return string(pass64)
}

// 微信消息解密
func WxMsgAesDecrypt(cipherText string, aesKey string) string {
	aesKeyByte := []byte(aesKey)
	bytesPass, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}
	tpass, err := aesDecrypt(bytesPass, aesKeyByte)
	if err != nil {
		panic(err)
	}
	util.GetLogger().Infof("解密后:%s\n", tpass)
	return string(tpass)
}
