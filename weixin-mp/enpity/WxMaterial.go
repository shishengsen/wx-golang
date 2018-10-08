package enpity

import "encoding/json"

type TempMaterial struct {
	Type			string				`json:"type"`
	MediaId			string				`json:"media_id"`
	CreateAt		string				`json:"create_at"`
}

type NewsMaterial struct {
	Articles			[]NewsArticle		`json:"articles"`
}

type NewsArticle struct {
	Title				string				`json:"title"`
	ThumbMediaId		string				`json:"thumb_media_id"`
	Author				string				`json:"author"`
	Digest				string				`json:"digest"`
	ShowCoverPic		int32				`json:"show_cover_pic"`
	Content				string				`json:"content"`
	ContentSourceUrl	string				`json:"content_source_url"`
}

//
type PermanentMaterial struct {
	MediaId				string				`json:"media_id"`
	Url					string				`json:"url"`
}

//
type VideoMaterialDesc struct {
	Title				string				`json:"title"`
	Introduction		string				`json:"introduction"`
}

type VideoPermanentMaterial struct {
	Title				string				`json:"title"`
	Description			string				`json:"description"`
	DownUrl				string				`json:"down_url"`
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