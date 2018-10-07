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
)

// post文本数据提交
func Post(url string, body string) ([]byte) {
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return result
}

// post文件数据提交
func PostWithFile(url string, file *os.File) ([]byte) {
	return PostWithFileAndBody(url, "", file)
}

func PostWithFileAndBody(url, body string, file *os.File) ([]byte) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	part, err := bodyWriter.CreateFormFile("file", file.Name())
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	if body != "" {
		var params map[string]string
		err = json.Unmarshal([]byte(body), &params)
		if err != nil {
			panic(err)
		}
		for key, val := range params {
			bodyWriter.WriteField(key, val)
		}
	}
	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.Close()
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return result
}

// get接口使用
func Get(url string) ([]byte) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return result
}
