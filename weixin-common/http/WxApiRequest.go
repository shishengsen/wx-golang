package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(url string, body string) {
	
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