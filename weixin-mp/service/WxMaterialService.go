package service

import (
	"encoding/json"
	"fmt"
	"os"
	"weixin-golang/weixin-common/http"
	"weixin-golang/weixin-common/wxconsts"
	"weixin-golang/weixin-mp/enpity"
)

const (
	add_temp_material			=			"https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	get_temp_material			=			"https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
	add_news_material			=			"https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%s"
	upload_news_pic				=			"https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"
	get_material_from_js		=			"https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=%s&media_id=%s"
	upload_other_material		=			"https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s"
	delete_material				=			"https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s"
)

// 增加临时素材
func (w *WeChat)AddTempMaterial(materialType string, file os.File) enpity.TempMaterial {
	reqUrl := fmt.Sprintf(add_temp_material, w.GetAccessToken(), materialType)
	msg := http.PostWithFile(reqUrl, &file)
	var material enpity.TempMaterial
	json.Unmarshal(msg, &material)
	return material
}

//TODO 返回byte[]数组暂未处理
func (w *WeChat)GetTempMaterial(materialId string) {
	reqUrl := fmt.Sprintf(get_temp_material,  w.GetAccessToken(), materialId)
	_ := http.Get(reqUrl)

}

//TODO 返回byte[]数组暂未处理
func (w *WeChat)GetHighSpeexMaterial(materialId string) {
	reqUrl := fmt.Sprintf(get_material_from_js, w.GetAccessToken(), materialId)
	_ := http.Get(reqUrl)
}

// 上传永久素材——新增永久图文素材
func (w *WeChat)AddNewsMaterial(newsMaterial enpity.NewsMaterial) enpity.TempMaterial {
	reqUrl := fmt.Sprintf(add_news_material, w.GetAccessToken())
	reqBody := newsMaterial.ToJson(newsMaterial)
	msg := http.Post(reqUrl, reqBody)
	var material enpity.TempMaterial
	json.Unmarshal(msg, &material)
	return material
}

// 上传图文消息内的图片获取URL
func (w *WeChat)UploadNewsMaterialPic(file os.File) string {
	reqUrl := fmt.Sprintf(upload_news_pic, w.GetAccessToken())
	msg := http.PostWithFile(reqUrl, &file)
	var tmp map[string]string
	json.Unmarshal(msg, &tmp)
	return tmp["url"]
}

// 新增其他类型永久素材
func (w *WeChat)AddNoVideoPermanentMaterial(materialType string, file os.File) enpity.PermanentMaterial {
	reqUrl := fmt.Sprintf(upload_other_material, w.GetAccessToken(), materialType)
	msg := http.PostWithFile(reqUrl, &file)
	var result enpity.PermanentMaterial
	json.Unmarshal(msg, &result)
	return result
}

// 新增视频类素材
func (w *WeChat)AddVideoMaterial(file os.File, desc enpity.VideoMaterialDesc) enpity.PermanentMaterial {
	reqUrl := fmt.Sprintf(upload_other_material, w.GetAccessToken(), wxconsts.TEMP_MATERIAL_VIDEO)
	body := desc.ToJson(desc)
	msg := http.PostWithFileAndBody(reqUrl, body, &file)

	var result enpity.PermanentMaterial
	json.Unmarshal(msg, &result)
	return result
}

func (w *WeChat)DeleteMaterial(mediaId string) {
	
}