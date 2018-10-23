package enpity

// 获取用户增减数据接口
type WxUserSummary struct {
	List				[]UserSummary			`json:"list"`
}

type UserSummary struct {
	RefDate				string					`json:"ref_date"`
	UserSource			int64					`json:"user_source"`
	NewUser				int64					`json:"new_user"`
	CancelUser			int64					`json:"cancel_user"`
}

// 获取累计用户数据接口
type WxUserCumulate struct {
	List				[]UserCumulate			`json:"list"`
}

type UserCumulate struct {
	RefDate				string					`json:"ref_date"`
	CumulateUser		int64					`json:"cumulate_user"`

}

// 获取图文群发每日数据接口
type WxArticleSummary struct {
	List				[]ArticleSummary		`json:"list"`
}

type ArticleSummary struct {
	RefDate				string					`json:"ref_date"`
	Msgid				string					`json:"msgid"`
	Title				string					`json:"title"`
	IntPageReadUser		int64					`json:"int_page_read_user"`
	IntPageReadCount	int64					`json:"int_page_read_count"`
	OriPageReadUser		int64					`json:"ori_page_read_user"`
	OriPageReadCount	int64					`json:"ori_page_read_count"`
	ShareUser			int64					`json:"share_user"`
	ShareCount			int64					`json:"share_count"`
	AddToFavUser		int64					`json:"add_to_fav_user"`
	AddToFavCount		int64					`json:"add_to_fav_count"`
}

// 获取图文群发总数据
type WxArticleTotal struct {
	RefDate				string					`json:"ref_date"`
	Msgid				string					`json:"msgid"`
	Title				string					`json:"title"`
	Details				[]ArticleTotalDetail	`json:"details"`
}

type ArticleTotalDetail struct {
	StatDate			string					`json:"stat_date"`
	TargetUser			int64					`json:"target_user"`
	IntPageReadUser		int64					`json:"int_page_read_user"`
	IntPageReadCount	int64					`json:"int_page_read_count"`
	OriPageReadUser		int64					`json:"ori_page_read_user"`
	OriPageReadCount	int64					`json:"ori_page_read_count"`
	ShareUser			int64					`json:"share_user"`
	ShareCount			int64					`json:"share_count"`
	AddToFavUser		int64					`json:"add_to_fav_user"`
	AddToFavCount		int64					`json:"add_to_fav_count"`
	IntPageFromSessionReadUser	int64			`json:"int_page_from_session_read_user"`
	IntPageFromSessionReadCount	int64			`json:"int_page_from_session_read_count"`
	IntPageFromHistMsgReadUser	int64			`json:"int_page_from_hist_msg_read_user"`
	IntPageFromHistMsgReadCount	int64			`json:"int_page_from_hist_msg_read_count"`
	IntPageFromFeedReadUser		int64			`json:"int_page_from_feed_read_user"`
	IntPageFromFeedReadCount	int64			`json:"int_page_from_feed_read_count"`
	IntPageFromFriendsReadUser	int64			`json:"int_page_from_friends_read_user"`
	IntPageFromFriendsReadCount int64			`json:"int_page_from_friends_read_count"`
	IntPageFromOtherReadUser	int64			`json:"int_page_from_other_read_user"`
	IntPageFromOtherReadCount	int64			`json:"int_page_from_other_read_count"`
	FeedShareFromSessionUser	int64			`json:"feed_share_from_session_user"`
	FeedShareFromSessionCnt		int64			`json:"feed_share_from_session_cnt"`
	FeedShareFromFeedUser		int64			`json:"feed_share_from_feed_user"`
	FeedShareFromFeedCnt		int64			`json:"feed_share_from_feed_cnt"`
	FeedShareFromOtherUser		int64			`json:"feed_share_from_other_user"`
	FeedShareFromOtherCnt		int64			`json:"feed_share_from_other_cnt"`
}

// 获取图文统计数据
type WxUserRead struct {
	List				[]UserRead				`json:"list"`
}

// 获取图文统计分时数据
type WxUserReadHour struct {
	List				[]UserRead			`json:"list"`
}

type UserRead struct {
	RefDate				string					`json:"ref_date"`
	RefHour				int64					`json:"ref_hour"`
	UserSource			int64					`json:"user_source"`
	IntPageReadUser		int64					`json:"int_page_read_user"`
	IntPageReadCount	int64					`json:"int_page_read_count"`
	OriPageReadUser		int64					`json:"ori_page_read_user"`
	OriPageReadCount	int64					`json:"ori_page_read_count"`
	ShareUser			int64					`json:"share_user"`
	ShareCount			int64					`json:"share_count"`
	AddToFavUser		int64					`json:"add_to_fav_user"`
	AddToFavCount		int64					`json:"add_to_fav_count"`
}


// 获取图文分享转发数据
type WxUserShare struct {
	List				[]UserShare			`json:"list"`
}

// 获取图文分享转发每日数据
type WxUserShareHour struct {
	List				[]UserShare		`json:"list"`
}

type UserShare struct {
	RefDate				string				`json:"ref_date"`
	RefHour				string				`json:"ref_hour"`
	ShareScene			string				`json:"share_scene"`
	ShareCount			string				`json:"share_count"`
	ShareUser			string				`json:"share_user"`
}

// 获取消息发送概况数据
type WxUpStreamMsg struct {
	List				[]UpStreamMsg		`json:"list"`
}

// 获取消息分送分时数据
type WxUpStreamMsgHour struct {
	List				[]UpStreamMsg	`json:"list"`
}

// 获取消息发送周数据
type WxUpStreamMsgWeek struct {
	List				[]UpStreamMsg		`json:"list"`
}

// 获取消息发送月数据
type WxUpStreamMsgMonth struct {
	List				[]UpStreamMsg		`json:"list"`
}

type UpStreamMsg struct {
	RefDate				string				`json:"ref_date"`
	RefHour				int64				`json:"ref_hour"`
	MsgType				int64				`json:"msg_type"`
	MsgUser				int64				`json:"msg_user"`
	MsgCount			int64				`json:"msg_count"`
}


// 获取消息发送分布数据
type WxUpStreamMsgDist struct {
	List				[]UpStreamMsgDist	`json:"list"`
}

// 获取消息发送分布周数据
type WxUpStreamMsgDistWeek struct {
	List				[]UpStreamMsgDist	`json:"list"`
}

// 获取消息发送分布月数据
type WxUpStreamMsgDistMonth struct {
	List				[]UpStreamMsgDist	`json:"list"`
}

type UpStreamMsgDist struct {
	RefDate				string				`json:"ref_date"`
	CountInterval		int64				`json:"count_interval"`
	MsgUser				int64				`json:"msg_user"`
}

// 获取接口分析数据
type WxInterfaceSummary struct {
	List				[]InterfaceSummary	`json:"list"`
}

// 获取接口分析分时数据
type WxInterfaceSummaryHour struct {
	List				[]InterfaceSummary	`json:"list"`
}

type InterfaceSummary struct {
	RefDate				string				`json:"ref_date"`
	RefHour				int64				`json:"ref_hour"`
	CallbackCount		int64				`json:"callback_count"`
	FailCount			int64				`json:"fail_count"`
	TotalTimeCost		int64				`json:"total_time_cost"`
	MaxTimeCost			int64				`json:"max_time_cost"`
}