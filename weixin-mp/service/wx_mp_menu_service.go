package service

import (
	"encoding/json"
	"fmt"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/utils"
	"wx-golang/weixin-mp/enpity"
)

const (
	/**
	创建自定义菜单
	 */
	create_menu = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
	/**
	查询自定义菜单
	 */
	query_menu  = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s"
	/**
	删除自定义菜单
	 */
	delete_menu = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s"
	/**
	创建个性化菜单
	 */
	create_selfdom_menu = "https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=%s"
	/**
	删除个性化菜单
	 */
	delete_selfdom_menu = "https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=%s"
	/**
	测试个性化菜单匹配结果
	 */
	test_selfdom_menu = "https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=%s"
)

// 自定义菜单创建接口
func (w *WeChat) WxMpCreateMenu(wxMenu enpity.WxMenu) {
	menuJson := wxMenu.ToJson(wxMenu)
	w.WxMpCreateMenuByJson(menuJson)
}

// 自定义菜单创建接口（前端定义的菜单json串，而不是在代码内部固定的）
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

//  个性化菜单创建接口
func (w *WeChat) WxCreateSelfDomMenu(wxMenu enpity.WxMenu) string {
	wxMenu.Matchrule.Language = w.cfg.Lang
	reqUrl := fmt.Sprintf(create_selfdom_menu, w.WxGetAccessToken())
	msg := createMenu(reqUrl, wxMenu.ToJson(wxMenu))
	var result map[string]string
	err := json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result["menuid"]
}

// 删除自定义菜单接口
func (w *WeChat) WxDeleteSelfDomMenu(menuId string) {
	reqUrl := fmt.Sprintf(delete_selfdom_menu, w.WxGetAccessToken())
	body := map[string]string {
		"menuid": menuId,
	}
	http.Post(reqUrl, string(utils.Interface2byte(body)))
}

func (w *WeChat) WxMatchSelfDomMenu(userId string) enpity.WxButton {
	reqUrl := fmt.Sprintf(test_selfdom_menu, w.WxGetAccessToken())
	body := map[string]string {
		"user_id": userId,
	}
	msg := http.Post(reqUrl, string(utils.Interface2byte(body)))
	var button enpity.WxButton
	err := json.Unmarshal(msg, &button)
	if err != nil {
		panic(err)
	}
	return button
}

// 创建自定义菜单的微信实际接口调用
func createMenu(reqUrl, body string) []byte {
	msg := http.Post(reqUrl, body)
	return msg
}

