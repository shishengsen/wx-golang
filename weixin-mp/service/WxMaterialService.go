package service

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/wxconsts"
	"wx-golang/weixin-mp/enpity"
	wxerr "wx-golang/weixin-common/error"
)

const (
	/**
	新增临时素材
	 */
	add_temp_material			=			"https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	/**
	获取临时素材
	 */
	get_temp_material			=			"https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
	/**
	新增永久图文素材上传图文素材
	 */
	add_news_material			=			"https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%s"
	/**
	上传图文消息内的图片获取URL
	 */
	upload_news_pic				=			"https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"
	/**
	高清语音素材获取接口
	 */
	get_material_from_js		=			"https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=%s&media_id=%s"
	/**
	新增其他类型永久素材
	 */
	add_other_material     		=			"https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s"
	/**
	获取永久素材
	 */
	get_permanent_material 		=			"https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s"
	/**
	删除永久素材
	 */
	delete_material        		=			"https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s"
)

// 增加临时素材
func (w *WeChat)AddTempMaterial(materialType string, file os.File) enpity.TempMaterial {
	reqUrl := fmt.Sprintf(add_temp_material, w.GetAccessToken(), materialType)
	msg := http.PostWithFile(reqUrl, &file)
	wxerr.WxMpErrorFromByte(msg, nil)
	var material enpity.TempMaterial
	json.Unmarshal(msg, &material)
	return material
}

//获取非视频微信临时素材
func (w *WeChat) GetTempNoVideoMaterial(materialId string) *os.File {
	reqUrl := fmt.Sprintf(get_temp_material,  w.GetAccessToken(), materialId)
	fileByte := http.Get(reqUrl)
	wxerr.WxMpErrorFromByte(fileByte, nil)
	_uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	videoFile, err := os.Create(_uuid.String())
	if err != nil {
		panic(err)
	}
	_, err = videoFile.Write(fileByte)
	if err != nil {
		panic(err)
	}
	return videoFile
}

// 获取微信视频临时素材
func (w *WeChat) GetTempVideoMaterial(materialId string) string {
	reqUrl := fmt.Sprintf(get_temp_material,  w.GetAccessToken(), materialId)
	result := http.Get(reqUrl)
	wxerr.WxMpErrorFromByte(result, nil)
	var tmp map[string]string
	err := json.Unmarshal(result, &tmp)
	if err != nil {
		panic(err)
	}
	if val, isPresent := tmp["video_url"]; isPresent == true {
		return val
	}
	panic("Error:[未找到 video_url]")
}

//高清语音素材获取接口
func (w *WeChat)GetHighSpeexMaterial(materialId string) *os.File {
	reqUrl := fmt.Sprintf(get_material_from_js, w.GetAccessToken(), materialId)
	voiceByte := http.Get(reqUrl)
	wxerr.WxMpErrorFromByte(voiceByte, nil)
	voiceFile, err := os.Create(materialId)
	if err != nil {
		panic(err)
	}
	_, err = voiceFile.Write(voiceByte)
	if err != nil {
		panic(err)
	}
	return voiceFile
}

// 上传永久素材——新增永久图文素材
func (w *WeChat)AddNewsMaterial(newsMaterial enpity.NewsMaterial) enpity.TempMaterial {
	reqUrl := fmt.Sprintf(add_news_material, w.GetAccessToken())
	reqBody := newsMaterial.ToJson(newsMaterial)
	msg := http.Post(reqUrl, reqBody)
	wxerr.WxMpErrorFromByte(msg, nil)
	var material enpity.TempMaterial
	json.Unmarshal(msg, &material)
	return material
}

// 上传图文消息内的图片获取URL
func (w *WeChat)UploadNewsMaterialPic(file os.File) string {
	reqUrl := fmt.Sprintf(upload_news_pic, w.GetAccessToken())
	msg := http.PostWithFile(reqUrl, &file)
	wxerr.WxMpErrorFromByte(msg, nil)
	var tmp map[string]string
	json.Unmarshal(msg, &tmp)
	return tmp["url"]
}

// 新增其他类型永久素材
func (w *WeChat)AddNoVideoPermanentMaterial(materialType string, file os.File) enpity.PermanentMaterial {
	reqUrl := fmt.Sprintf(add_other_material, w.GetAccessToken(), materialType)
	msg := http.PostWithFile(reqUrl, &file)
	wxerr.WxMpErrorFromByte(msg, nil)
	var result enpity.PermanentMaterial
	json.Unmarshal(msg, &result)
	return result
}

// 新增视频类素材
func (w *WeChat)AddVideoMaterial(file os.File, desc enpity.VideoMaterialDesc) enpity.PermanentMaterial {
	reqUrl := fmt.Sprintf(add_other_material, w.GetAccessToken(), wxconsts.TEMP_MATERIAL_VIDEO)
	body := desc.ToJson(desc)
	msg := http.PostWithFileAndBody(reqUrl, body, &file)
	wxerr.WxMpErrorFromByte(msg, nil)
	var result enpity.PermanentMaterial
	json.Unmarshal(msg, &result)
	return result
}

// 获取图文永久素材
func (w *WeChat)GetPermanentNewsMaterial(materialId string) enpity.NewsMaterial {
	result := getPermanentMaterial(materialId, w.GetAccessToken())
	var newsMaterial enpity.NewsMaterial
	err := json.Unmarshal(result, &newsMaterial)
	if err != nil {
		panic(err)
	}
	return newsMaterial
}

// 获取视频永久素材
func (w *WeChat)GetPermanentVideoMaterial(materialId string) enpity.VideoPermanentMaterial {
	result := getPermanentMaterial(materialId, w.GetAccessToken())
	var videoMaterial enpity.VideoPermanentMaterial
	err := json.Unmarshal(result, &videoMaterial)
	if err != nil {
		panic(err)
	}
	return videoMaterial
}

// 获取其他永久素材
func (w *WeChat)GetOtherPermanentMaterial(materialId string) *os.File {
	result := getPermanentMaterial(materialId, w.GetAccessToken())
	materialFile, err := os.Create(materialId)
	if err != nil {
		panic(err)
	}
	_, err = materialFile.Write(result)
	if err != nil {
		panic(err)
	}
	return materialFile
}

func getPermanentMaterial(materialId, accessToken string) []byte {
	reqUrl := fmt.Sprintf(get_permanent_material, accessToken)
	body := map[string]string {
		"media_id": materialId,
	}
	bodyByte, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	result := http.Post(reqUrl, string(bodyByte))
	wxerr.WxMpErrorFromByte(result, nil)
	return result
}

// 删除永久素材
func (w *WeChat)DeleteMaterial(mediaId string) bool {
	reqUrl := fmt.Sprintf(get_permanent_material, w.GetAccessToken())
	body := map[string]string {
		"media_id": mediaId,
	}
	bodyByte, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	result := http.Post(reqUrl, string(bodyByte))
	wxerr.WxMpErrorFromByte(result, nil)
	return true
}