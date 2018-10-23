package enpity

import "encoding/json"

//
type TempMaterial struct {
	Type			string				`json:"type"`
	MediaId			string				`json:"media_id"`
	CreateAt		string				`json:"create_at"`
}

//
type NewsMaterial struct {
	Articles			[]NewsArticle		`json:"articles"`
}

// 永久图文素材
type NewsArticle struct {
	Title				string				`json:"title"`
	ThumbMediaId		string				`json:"thumb_media_id"`
	Author				string				`json:"author"`
	Digest				string				`json:"digest"`
	ShowCoverPic		int32				`json:"show_cover_pic"`
	Content				string				`json:"content"`
	ContentSourceUrl	string				`json:"content_source_url"`
	NeedOpenComment		int32				`json:"need_open_comment,omitempty"`
	OnlyFansCanComment	int32				`json:"only_fans_can_comment,omitempty"`
}

// 新增素材返回的对象信息
type PermanentMaterial struct {
	MediaId				string				`json:"media_id"`
	Url					string				`json:"url"`
}

// 在上传视频素材时需要POST另一个表单，id为description，包含素材的描述信息，内容格式为JSON
type VideoMaterialDesc struct {
	Title				string				`json:"title"`
	Introduction		string				`json:"introduction"`
}

// 视频消息素材
type VideoPermanentMaterial struct {
	Title				string				`json:"title"`
	Description			string				`json:"description"`
	DownUrl				string				`json:"down_url"`
}

// 获取素材总数
type MaterialTotalNum struct {
	VoiceCount			int64				`json:"voice_count"`
	VideoCount			int64				`json:"video_count"`
	ImageCount			int64				`json:"image_count"`
	NewsCount			int64				`json:"news_count"`
}

// 素材列表
type NewsMaterialList struct {
	TotalCount			int64				`json:"total_count"`
	ItemCount			int64				`json:"item_count"`
}

// 永久图文消息素材列表
type NewsMaterialItem struct {
	MediaId				string				`json:"media_id"`
	Content				struct{
		NewsItem		[]NewsArticle		`json:"news_item"`
	}
	UpdateTime			string				`json:"update_time"`
}

// 其他类型（图片、语音、视频）素材列表
type OtherMaterialList struct {
	TotalCount			int64				`json:"total_count"`
	ItemCount			int64				`json:"item_count"`
	Item				[]struct{
		MediaId			string				`json:"media_id"`
		Name			string				`json:"name"`
		UpdateTime		string				`json:"update_time"`
		Url				string				`json:"url"`
	}
}

func (n *NewsMaterial)ToJson(material NewsMaterial) string {
	newsJson, err := json.Marshal(material)
	if err != nil {
		panic(err)
	}
	return string(newsJson)
}

func (d *VideoMaterialDesc)ToJson(desc VideoMaterialDesc) string {
	descJson, err := json.Marshal(desc)
	if err != nil {
		panic(err)
	}
	return string(descJson)
}