package service

import (
	"encoding/json"
	"fmt"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-mp/enpity"
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
	reqUrl := fmt.Sprintf(create_menu, w.WxGetAccessToken())
	createMenu(reqUrl, menuJson)
}

// 自定义菜单查询接口
func (w *WeChat) WxMpQueryMenu() enpity.WxMenu {
	reqUrl := fmt.Sprintf(query_menu, w.WxGetAccessToken())
	msg := http.Get(reqUrl)
	var menu enpity.WxMenu
	json.Unmarshal(msg, &menu)
	return menu
}

// 自定义菜单删除接口
func (w *WeChat) WxMpDeleteMenu() string {
	reqUrl := fmt.Sprintf(delete_menu, w.WxGetAccessToken())
	msg := http.Get(reqUrl)
	return string(msg)
}

// 创建自定义菜单
func createMenu(reqUrl, body string) string {
	msg := http.Post(reqUrl, body)
	return string(msg)
}
