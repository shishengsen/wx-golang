package enpity

// 微信消息推送消息结构体
type WxMessage struct {
	ToUserName   string `xml:"ToUserName,omitempty"`   // 开发者微信号
	FromUserName string `xml:"FromUserName,omitempty"` // 发送方帐号（一个OpenID）
	CreateTime   int64  `xml:"CreateTime,omitempty"`   // 消息创建时间 （整型）
	MsgType      string `xml:"MsgType,omitempty"`      // 消息类型
	MediaId      string `xml:"MediaId,omitempty"`      // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	MsgId        int64  `xml:"MsgId,omitempty"`        // 消息id，64位整型

	// 文本消息
	Content string `xml:"Content,omitempty"` // 文本消息内容

	// 图片消息
	PicUrl string `xml:"PicUrl,omitempty"` // 图片链接（由系统生成）

	// 语音消息
	Format      string `xml:"Format,omitempty"`      // 语音格式：amr
	Recognition string `xml:"Recongition,omitempty"` // 语音识别结果，UTF8编码

	// 视频消息
	ThumbMediaId string `xml:"ThumbMediaid,omitempty"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。

	// 地理位置消息
	LocationX float64 `xml:"Location_X,omitempty"` // 地理位置维度
	LocationY float64 `xml:"Location_Y,omitempty"` // 地理位置经度
	Scale     int32   `xml:"Scale,omitempty"`      // 地图缩放大小
	Label     string  `xml:"Label,omitempty"`      // 地理位置信息

	// 链接消息
	Title       string `xml:"Title,omitempty"`       // 消息标题
	Description string `xml:"Description,omitempty"` // 消息描述
	Url         string `xml:"Url,omitempty"`         // 消息链接

	// 关注、取消事件
	Event    string `xml:"Event,omitempty"`    // 事件类型
	EventKey string `xml:"EventKey,omitempty"` // 事件KEY值
	Ticket   string `xml:"Ticket,omitempty"`   // 二维码的ticket，可用来换取二维码图片

	// 上报地理位置事件
	Latitude  float64 `xml:"Latitude,omitempty"`  // 地理位置纬度
	Longitude float64 `xml:"Longitude,omitempty"` // 地理位置经度
	Precision string  `xml:"Precision,omitempty"` // 地理位置精度

}
