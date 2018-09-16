package http

import (
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
)

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