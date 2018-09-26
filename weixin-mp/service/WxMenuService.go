package service

import (
	"encoding/json"
	"fmt"
	"weixin-golang/weixin-common/http"
	"weixin-golang/weixin-mp/enpity"
)

const (
	create_menu = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
	query_menu  = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s"
	delete_menu = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s"
)

// 自定义菜单创建接口
func (w *WeChat) WxMpCreateMenu(wxMenu enpity.WxMenu) {
	menuJson := wxMenu.ToJson(wxMenu)
	w.WxMpCreateMenuByJson(menuJson)
}

// 自定义菜单创建接口()
func (w *WeChat) WxMpCreateMenuByJson(menuJson string) {
	reqUrl := fmt.Sprintf(create_menu, w.GetAccessToken())
	createMenu(reqUrl, menuJson)
}

// 自定义菜单查询接口
func (w *WeChat) WxMpQueryMenu() enpity.WxMenu {
	reqUrl := fmt.Sprintf(query_menu, w.GetAccessToken())
	msg, err := http.Get(reqUrl)
	if err != nil {
		panic(err)
	}
	var menu enpity.WxMenu
	json.Unmarshal(msg, &menu)
	return menu
}

// 自定义菜单删除接口
func (w *WeChat) WxMpDeleteMenu() string {
	reqUrl := fmt.Sprintf(delete_menu, w.GetAccessToken())
	msg, err := http.Get(reqUrl)
	if err != nil {
		panic(err)
	}
	return string(msg)
}

// 创建自定义菜单
func createMenu(reqUrl, body string) string {
	msg, err := http.Post(reqUrl, body)
	if err != nil {
		panic(err)
	}
	return string(msg)
}
