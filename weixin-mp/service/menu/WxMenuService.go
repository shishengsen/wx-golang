package menu

import (
	"fmt"
	"encoding/json"
	"weixin-golang/weixin-mp/enpity"
	"weixin-golang/weixin-mp/service"
	"weixin-golang/weixin-common/http"
)

const (
	create_menu			=			"https://api.weixin.qq.com/cgi-bin/menu/create?access_token="
	query_menu			=			"https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s"
)

// 自定义菜单创建接口
func WxMpCreateMenu(wxMenu enpity.WxMenu) {
	menuJson := wxMenu.ToJson(wxMenu)
	WxMpCreateMenuByJson(menuJson)
}

// 自定义菜单创建接口()
func WxMpCreateMenuByJson(menuJson string) {
	createMenu(menuJson)
}

// 自定义菜单查询接口
func WxMpQueryMenu() enpity.WxMenu {
	req_url := fmt.Sprintf(query_menu, service.GetAccessToken())
	msg, err := http.Get(req_url)
	if err != nil {
		panic(err)
	}
	var menu enpity.WxMenu
	json.Unmarshal(msg, &menu)
	return menu
}

// 自定义菜单删除接口
func WxMpDeleteMenu() string {
	return ""
}

// 创建自定义菜单
func createMenu(body string) string {
	req_url := create_menu + service.GetAccessToken()
	msg, err := http.Post(req_url, body)
	if err != nil {
		panic(err)
	}
	return string(msg)
}