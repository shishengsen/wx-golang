package enpity

import (
	"encoding/json"
)

type WxKf struct {
	KfAccount string `json:"kf_account,omitempty"`
	NickName  string `json:"nickname,omitempty"`
	Password  string `json:"password,omitempty"`
}

type WxKfs struct {
	Kfs []WxKf `json:"kf_list,omitempty"`
}

type WxKfMsg struct {
	ToUser  string        `json:"touser"`
	MsgType string        `json:"msgtype"`
	Text    KfTextMsg     `json:"text,omitempty"`
	Image   KfMaterialMsg `json:"image,omitempty"`
	Voice   KfMaterialMsg `json:"voice,omitempty"`
	MpNews  KfMaterialMsg `json:"mpnews,omitempty"`
	Video   KfVideoMsg    `json:"video,omitempty"`
	Music   KfMusicMsg    `json:"music,omitempty"`
	WxCard  KfCard        `json:"wxcard,omitempty"`
}

type KfTextMsg struct {
	Content string `json:"content"`
}

type KfMaterialMsg struct {
	MediaId string `json:"media_id"`
}

type KfVideoMsg struct {
	MediaId      string `json:"media_id"`
	ThumbMediaId string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

type KfMusicMsg struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicUrl     string `json:"musicurl"`
	HqMusicUrl   string `json:"hqmusicurl"`
	ThumbMediaId string `json:"thumbmediaid"`
}

type KfNews struct {
	Articles []Article `json:"articles"`
}

type KfCard struct {
	CardId string `json:"card_id"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

func (this *WxKf) ToJson(kf WxKf) string {
	kfJson, err := json.Marshal(kf)
	if err != nil {
		panic(err)
	}
	return string(kfJson)
}

func (this *WxKfMsg) ToJson(kfMsg WxKfMsg) string {
	kfMsgJson, err := json.Marshal(kfMsg)
	if err != nil {
		panic(err)
	}
	return string(kfMsgJson)
}
