package service

import (
	"fmt"
	"strings"
	"time"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/utils"
	wxerr "wx-golang/weixin-common/error"
)

const (
	/**
	获取用户增减数据（getusersummary）
	 */
	get_user_summary			=				"https://api.weixin.qq.com/datacube/getusersummary?access_token=%s"
	/**
	获取累计用户数据（getusercumulate）
	 */
	get_user_cumulate			=				"https://api.weixin.qq.com/datacube/getusercumulate?access_token=%s"
)

func (w *WeChat) WxMpGetUserSummary(start, end time.Time) {
	startS := strings.Split(start.Format("2006-01-02 15:04:05"), " ")[0]
	endS := strings.Split(end.Format("2006-01-02 15:04:05"), " ")[0]
	reqUrl := fmt.Sprintf(get_user_summary, w.WxGetAccessToken())
	body := map[string]string{
		"begin_date": startS,
		"end_date": endS,
	}
	msg := http.Post(reqUrl, string(utils.Interface2byte(body)))
	wxerr.WxMpErrorFromByte(msg, nil)
}
