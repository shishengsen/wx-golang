package service

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"wx-golang/weixin-common/http"
	"wx-golang/weixin-common/utils"
	"wx-golang/weixin-common/wxconsts"
	"wx-golang/weixin-mp/enpity"
)

const (
	/**
	新增临时素材
	*/
	add_temp_material = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	/**
	获取临时素材
	*/
	get_temp_material = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
	/**
	新增永久图文素材上传图文素材
	*/
	add_news_material = "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%s"
	/**
	更新图文素材
	*/
	update_news_material = "https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=%s"
	/**
	上传图文消息内的图片获取URL
	*/
	upload_news_pic = "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"
	/**
	高清语音素材获取接口
	*/
	get_material_from_js = "https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=%s&media_id=%s"
	/**
	新增其他类型永久素材
	*/
	add_other_material = "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s"
	/**
	获取永久素材
	*/
	get_permanent_material = "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s"
	/**
	删除永久素材
	*/
	delete_material = "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s"
	/**
	获取素材总数
	*/
	get_material_total_nums = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=%s"
	/**
	获取素材列表
	*/
	get_material_list = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s"
	/**
	打开已群发文章评论
	*/
	get_comment_open = "https://api.weixin.qq.com/cgi-bin/comment/open?access_token=%s"
	/**
	关闭已群发文章评论
	*/
	close_comment = "https://api.weixin.qq.com/cgi-bin/comment/close?access_token=%s"
	/**
	查看指定文章的评论数据
	*/
	see_comment = "https://api.weixin.qq.com/cgi-bin/comment/list?access_token=%s"
	/**
	将评论标记精选
	*/
	mark_elect_comment = "https://api.weixin.qq.com/cgi-bin/comment/markelect?access_token=%s"
	/**
	取消评论精选
	*/
	cancle_elect_comment = "https://api.weixin.qq.com/cgi-bin/comment/unmarkelect?access_token=%s"
	/**
	删除评论
	*/
	delete_comment = "https://api.weixin.qq.com/cgi-bin/comment/delete?access_token=%s"
	/**
	回复评论
	*/
	reply_comment = "https://api.weixin.qq.com/cgi-bin/comment/reply/add?access_token=%s"
	/**
	删除回复
	*/
	delete_reply_comment = "https://api.weixin.qq.com/cgi-bin/comment/reply/delete?access_token=%s"
)

type WxMaterial struct {
	token	*Token
} 

// 增加临时素材
func (m *WxMaterial) WxAddTempMaterial(materialType string, file os.File) (enpity.TempMaterial, error) {
	reqUrl := fmt.Sprintf(add_temp_material, m.token.WxGetAccessToken(), materialType)
	msg, err := http.PostWithFile(reqUrl, &file)
	var material enpity.TempMaterial
	json.Unmarshal(msg, &material)
	return material, err
}

//获取非视频微信临时素材
func (m *WxMaterial) WxGetTempNoVideoMaterial(materialId string) *os.File {
	reqUrl := fmt.Sprintf(get_temp_material, m.token.WxGetAccessToken(), materialId)
	fileByte, err := http.Get(reqUrl)
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
func (m *WxMaterial) WxGetTempVideoMaterial(materialId string) string {
	reqUrl := fmt.Sprintf(get_temp_material, m.token.WxGetAccessToken(), materialId)
	result, err := http.Get(reqUrl)
	var tmp map[string]string
	err = json.Unmarshal(result, &tmp)
	if err != nil {
		panic(err)
	}
	if val, isPresent := tmp["video_url"]; isPresent == true {
		return val
	}
	panic("Error:[未找到 video_url]")
}

// 高清语音素材获取接口
func (m *WxMaterial) WxGetHighSpeexMaterial(materialId string) *os.File {
	reqUrl := fmt.Sprintf(get_material_from_js, m.token.WxGetAccessToken(), materialId)
	voiceByte, err := http.Get(reqUrl)
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
func (m *WxMaterial) WxAddNewsMaterial(newsMaterial enpity.NewsMaterial) (enpity.TempMaterial, error) {
	reqUrl := fmt.Sprintf(add_news_material, m.token.WxGetAccessToken())
	reqBody := newsMaterial.ToJson(newsMaterial)
	msg, err := http.Post(reqUrl, reqBody)
	var material enpity.TempMaterial
	json.Unmarshal(msg, &material)
	return material, err
}

// 对永久图文素材进行修改
func (m *WxMaterial) UpdateNewsMaterial(newsMaterial enpity.NewsMaterial) bool {
	reqUrl := fmt.Sprintf(update_news_material, m.token.WxGetAccessToken())
	reqBody := newsMaterial.ToJson(newsMaterial)
	http.Post(reqUrl, reqBody)
	return true
}

// 上传图文消息内的图片获取URL
func (m *WxMaterial) WxUploadNewsMaterialPic(file os.File) (string, error) {
	reqUrl := fmt.Sprintf(upload_news_pic, m.token.WxGetAccessToken())
	msg, err := http.PostWithFile(reqUrl, &file)
	var tmp map[string]string
	json.Unmarshal(msg, &tmp)
	return tmp["url"], err
}

// 新增其他类型永久素材
func (m *WxMaterial) WxAddNoVideoPermanentMaterial(materialType string, file os.File) (enpity.PermanentMaterial, error) {
	reqUrl := fmt.Sprintf(add_other_material, m.token.WxGetAccessToken(), materialType)
	msg, err := http.PostWithFile(reqUrl, &file)
	var result enpity.PermanentMaterial
	json.Unmarshal(msg, &result)
	return result, err
}

// 新增视频类素材
func (m *WxMaterial) WxAddVideoMaterial(file os.File, desc enpity.VideoMaterialDesc) (enpity.PermanentMaterial, error) {
	reqUrl := fmt.Sprintf(add_other_material, m.token.WxGetAccessToken(), wxconsts.TEMP_MATERIAL_VIDEO)
	body := desc.ToJson(desc)
	msg, err := http.PostWithFileAndBody(reqUrl, body, &file)
	var result enpity.PermanentMaterial
	json.Unmarshal(msg, &result)
	return result, err
}

// 获取图文永久素材
func (m *WxMaterial) WxGetPermanentNewsMaterial(materialId string) (enpity.NewsMaterial, error) {
	result, err := getPermanentMaterial(materialId, m.token.WxGetAccessToken())
	var newsMaterial enpity.NewsMaterial
	json.Unmarshal(result, &newsMaterial)
	return newsMaterial, err
}

// 获取视频永久素材
func (m *WxMaterial) WxGetPermanentVideoMaterial(materialId string) (enpity.VideoPermanentMaterial, error) {
	result, err := getPermanentMaterial(materialId, m.token.WxGetAccessToken())
	var videoMaterial enpity.VideoPermanentMaterial
	json.Unmarshal(result, &videoMaterial)
	return videoMaterial, err
}

// 获取其他永久素材
func (m *WxMaterial) WxGetOtherPermanentMaterial(materialId string) (*os.File, error) {
	result, err := getPermanentMaterial(materialId, m.token.WxGetAccessToken())
	materialFile, err := os.Create(materialId)
	_, err = materialFile.Write(result)
	return materialFile, err
}

// 删除永久素材
func (m *WxMaterial) WxDeleteMaterial(mediaId string) bool {
	reqUrl := fmt.Sprintf(delete_material, m.token.WxGetAccessToken())
	body := map[string]string{
		"media_id": mediaId,
	}
	http.Post(reqUrl, string(utils.Interface2byte(body)))
	return true
}

// 获取素材总数
func (m *WxMaterial) WxGetMaterialTotalNum() enpity.MaterialTotalNum {
	reqUrl := fmt.Sprintf(get_material_total_nums, m.token.WxGetAccessToken())
	msg, err := http.Get(reqUrl)
	var totalNum enpity.MaterialTotalNum
	err = json.Unmarshal(msg, &totalNum)
	if err != nil {
		panic(err)
	}
	return totalNum
}

// 获取永久图文消息素材列表
func (m *WxMaterial) WxGetNewsMaterialList(offset, count int64) enpity.NewsMaterialList {
	reqUrl := fmt.Sprintf(get_material_list, m.token.WxGetAccessToken())
	body := map[string]interface{}{
		"type":   wxconsts.MATERIAL_NEWS,
		"offset": offset,
		"count":  count,
	}
	msg, err := http.Post(reqUrl, string(utils.Interface2byte(body)))
	var result enpity.NewsMaterialList
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 获取其他类型（图片、语音、视频）素材列表
func (m *WxMaterial) WxGetOtherMaterialList(materialType string, offset, count int64) enpity.OtherMaterialList {
	reqUrl := fmt.Sprintf(get_material_list, m.token.WxGetAccessToken())
	body := map[string]interface{}{
		"type":   materialType,
		"offset": offset,
		"count":  count,
	}
	msg, err := http.Post(reqUrl, string(utils.Interface2byte(body)))
	var result enpity.OtherMaterialList
	err = json.Unmarshal(msg, &result)
	if err != nil {
		panic(err)
	}
	return result
}

//
func getPermanentMaterial(materialId, accessToken string) ([]byte, error) {
	reqUrl := fmt.Sprintf(get_permanent_material, accessToken)
	body := map[string]string{
		"media_id": materialId,
	}
	result, err := http.Post(reqUrl, string(utils.Interface2byte(body)))
	return result, err
}
