package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	wxErr "wx-golang/weixin-common/error"
	"wx-golang/weixin-common/log"
)

// post纯文本数据提交
func Post(url string, body string) ([]byte, error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, wxErr.WxMpErrorFromByte(result, resp)
}

// post纯媒体数据提交
func PostWithFile(url string, file *os.File) ([]byte, error) {
	return PostWithFileAndBody(url, "", file)
}

// post文件表单
func PostWithFileAndBody(url, body string, file *os.File) ([]byte, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	part, err := bodyWriter.CreateFormFile("file", file.Name())
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	if body != "" {
		var params map[string]string
		json.Unmarshal([]byte(body), &params)
		for key, val := range params {
			err = bodyWriter.WriteField(key, val)
			if err != nil {
				return nil, err
			}
		}
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer func() {
		resp.Body.Close()
		if r := recover(); r != nil {
			log.GetLogger().Info(r)
		}
	}()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, wxErr.WxMpErrorFromByte(result, resp)
}

// get请求
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {

	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, wxErr.WxMpErrorFromByte(result, resp)
}
