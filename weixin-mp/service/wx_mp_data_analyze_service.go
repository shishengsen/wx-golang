package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	wxerr "wx-golang/weixin-common/error"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/utils"
	"wx-golang/weixin-mp/enpity"
)

const (
	/**
	获取用户增减数据（getusersummary）
	*/
	get_user_summary = "https://api.weixin.qq.com/datacube/getusersummary?access_token=%s"
	/**
	获取累计用户数据（getusercumulate）
	*/
	get_user_cumulate = "https://api.weixin.qq.com/datacube/getusercumulate?access_token=%s"
	/**
	获取图文群发每日数据（getarticlesummary）
	*/
	get_article_summary = "https://api.weixin.qq.com/datacube/getarticlesummary?access_token=%s"
	/**
	获取图文群发总数据（getarticletotal）
	*/
	get_article_total = "https://api.weixin.qq.com/datacube/getarticletotal?access_token=%s"
	/**
	获取图文统计数据（getuserread）
	*/
	get_user_read = "https://api.weixin.qq.com/datacube/getuserread?access_token=%s"
	/**
	获取图文统计分时数据（getuserreadhour）
	*/
	get_user_read_hour = "https://api.weixin.qq.com/datacube/getuserreadhour?access_token=%s"
	/**
	获取图文分享转发数据（getusershare）
	*/
	get_user_share = "https://api.weixin.qq.com/datacube/getusershare?access_token=%s"
	/**
	获取图文分享转发分时数据（getusersharehour）
	*/
	get_user_share_hour = "https://api.weixin.qq.com/datacube/getusersharehour?access_token=%s"
	/**
	获取消息发送概况数据（getupstreammsg）
	*/
	get_up_stream_msg = "https://api.weixin.qq.com/datacube/getupstreammsg?access_token=%s"
	/**
	获取消息分送分时数据（getupstreammsghour）
	*/
	get_up_stream_msg_hour = "https://api.weixin.qq.com/datacube/getupstreammsghour?access_token=%s"
	/**
	获取消息发送周数据（getupstreammsgweek）
	*/
	get_up_stream_msg_week = "https://api.weixin.qq.com/datacube/getupstreammsgweek?access_token=%s"
	/**
	获取消息发送月数据（getupstreammsgmonth）
	*/
	get_up_stream_msg_month = "https://api.weixin.qq.com/datacube/getupstreammsgmonth?access_token=%s"
	/**
	获取消息发送分布数据（getupstreammsgdist）
	*/
	get_up_stream_msg_dist = "https://api.weixin.qq.com/datacube/getupstreammsgdist?access_token=%s"
	/**
	获取消息发送分布周数据（getupstreammsgdistweek）
	*/
	get_up_stream_msg_dist_week = "https://api.weixin.qq.com/datacube/getupstreammsgdistweek?access_token=%s"
	/**
	获取消息发送分布月数据（getupstreammsgdistmonth）
	*/
	get_up_stream_msg_dist_month = "https://api.weixin.qq.com/datacube/getupstreammsgdistmonth?access_token=%s"
	/**
	获取接口分析数据（getinterfacesummary）
	*/
	get_interface_summary = "https://api.weixin.qq.com/datacube/getinterfacesummary?access_token=%s"
	/**
	获取接口分析分时数据（getinterfacesummaryhour）
	*/
	get_interface_summary_hour = "https://api.weixin.qq.com/datacube/getinterfacesummaryhour?access_token=%s"
)

// 获取用户增减数据（getusersummary）
func (w *WeChat) WxMpGetUserSummary(start, end time.Time) enpity.WxUserSummary {
	reqUrl := fmt.Sprintf(get_user_summary, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxUserSummary
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取累计用户数据（getusercumulate）
func (w *WeChat) WxMpGetUserCumulate(start, end time.Time) enpity.WxUserCumulate {
	reqUrl := fmt.Sprintf(get_user_cumulate, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxUserCumulate
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取图文群发每日数据（getarticlesummary）
func (w *WeChat) WxMpGetArticleSummary(start, end time.Time) enpity.WxArticleSummary {
	reqUrl := fmt.Sprintf(get_article_summary, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxArticleSummary
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取图文群发总数据（getarticletotal）
func (w *WeChat) WxMpGetArticleTotal(start, end time.Time) enpity.WxArticleTotal {
	reqUrl := fmt.Sprintf(get_article_total, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxArticleTotal
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取图文统计数据（getuserread）
func (w *WeChat) WxMpGetUserRead(start, end time.Time) enpity.WxUserRead {
	reqUrl := fmt.Sprintf(get_user_read, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxUserRead
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取图文统计分时数据（getuserreadhour）
func (w *WeChat) WxMpGetUserReadHour(start, end time.Time) enpity.WxUserReadHour {
	reqUrl := fmt.Sprintf(get_user_read_hour, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxUserReadHour
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取图文分享转发数据（getusershare）
func (w *WeChat) WxMpGetUserShare(start, end time.Time) enpity.WxUserShare {
	reqUrl := fmt.Sprintf(get_user_share, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxUserShare
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取图文分享转发分时数据（getusersharehour）
func (w *WeChat) WxMpGetUserShareHour(start, end time.Time) enpity.WxUserShareHour {
	reqUrl := fmt.Sprintf(get_user_share_hour, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxUserShareHour
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取消息发送概况数据（getupstreammsg）
func (w *WeChat) WxMpGetUpStreamMsg(start, end time.Time) enpity.WxUpStreamMsg {
	reqUrl := fmt.Sprintf(get_up_stream_msg, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxUpStreamMsg
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取消息分送分时数据（getupstreammsghour）
func (w *WeChat) WxMpGetUpStreamMsgHour(start, end time.Time) enpity.WxUpStreamMsgHour {
	reqUrl := fmt.Sprintf(get_up_stream_msg_hour, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxUpStreamMsgHour
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取消息发送周数据（getupstreammsgweek）
func (w *WeChat) WxMpGetUpStreamMsgWeek(start, end time.Time) enpity.WxUpStreamMsgWeek {
	reqUrl := fmt.Sprintf(get_up_stream_msg_week, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxUpStreamMsgWeek
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取消息发送月数据（getupstreammsgmonth）
func (w *WeChat) WxMpGetUpStreamMsgMonth(start, end time.Time) enpity.WxUpStreamMsgMonth {
	reqUrl := fmt.Sprintf(get_up_stream_msg_month, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	wxerr.WxMpErrorFromByte(msg, nil)
	var result enpity.WxUpStreamMsgMonth
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取消息发送分布数据（getupstreammsgdist）
func (w *WeChat) WxMpGetUpStreamMsgDist(start, end time.Time) enpity.WxUpStreamMsgDist {
	reqUrl := fmt.Sprintf(get_up_stream_msg_dist, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	wxerr.WxMpErrorFromByte(msg, nil)
	var result enpity.WxUpStreamMsgDist
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取消息发送分布周数据（getupstreammsgdistweek）
func (w *WeChat) WxMpGetUpStreamMsgDistWeek(start, end time.Time) enpity.WxUpStreamMsgDistWeek {
	reqUrl := fmt.Sprintf(get_up_stream_msg_dist_week, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	wxerr.WxMpErrorFromByte(msg, nil)
	var result enpity.WxUpStreamMsgDistWeek
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取消息发送分布月数据（getupstreammsgdistmonth）
func (w *WeChat) WxMpGetUpStreamMsgDistMonth(start, end time.Time) enpity.WxUpStreamMsgDistMonth {
	reqUrl := fmt.Sprintf(get_up_stream_msg_dist_month, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxUpStreamMsgDistMonth
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取接口分析数据（getinterfacesummary）
func (w *WeChat) WxGetInterfaceSummary(start, end time.Time) enpity.WxInterfaceSummary {
	reqUrl := fmt.Sprintf(get_interface_summary, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxInterfaceSummary
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取接口分析分时数据（getinterfacesummaryhour）
func (w *WeChat) WxGetInterfaceSummaryHour(start, end time.Time) enpity.WxInterfaceSummaryHour {
	reqUrl := fmt.Sprintf(get_interface_summary_hour, w.WxGetAccessToken())
	msg, err := wxDataAnalyzeRequest(reqUrl, start, end)
	var result enpity.WxInterfaceSummaryHour
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func wxDataAnalyzeRequest(reqUrl string, start, end time.Time) ([]byte, error) {
	startS := strings.Split(utils.TimeFormatToString(start), " ")[0]
	endS := strings.Split(utils.TimeFormatToString(end), " ")[0]
	body := map[string]string{
		"begin_date": startS,
		"end_date":   endS,
	}
	return http.Post(reqUrl, string(utils.Interface2byte(body)))
}
