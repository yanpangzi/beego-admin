package controllers

import "beego-admin/services"

type EditorController struct {
	baseController
}

func (this *EditorController) Server() {
	var ueditorService services.UeditorService
	action := this.GetString("action")
	switch action {
	case "config":
		result := ueditorService.GetConfig()
		this.Data["json"] = result
		this.ServeJSON()
	case "uploadimage"://上传图片
	case "uploadscrawl"://上传涂鸦
	case "uploadvideo"://上传视频
	case "uploadfile"://上传文件
	case "listimage"://列出图片
	case "listfile"://列出文件
	case "catchimage"://抓取远程文件
	default:
		this.Data["json"] = map[string]string{
			"state":"请求地址出错",
		}
		this.ServeJSON()
	}

}
