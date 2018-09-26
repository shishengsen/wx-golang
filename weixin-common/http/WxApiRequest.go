package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

// post文本数据提交
func Post(url string, body string) ([]byte, error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", url, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

// post文件数据提交
func PostWithFile(url string, file multipart.File) ([]byte, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	part, err := bodyWriter.CreateFormFile("file", "tx.png")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, bodyBuf)
	if err != nil {
		panic(err)
	}
	contentType := bodyWriter.FormDataContentType()
	resp, err := http.Post(url, contentType, bodyBuf)
	return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", url, resp.StatusCode)
}

func Get(url string) ([]byte, error) {
	resp, err_resp := http.Get(url)
	if err_resp != nil {
		panic(err_resp)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", url, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
