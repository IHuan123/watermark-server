package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"watermarkServer/handleVideo"
	"watermarkServer/modules"
)

const (
	//抖音地址
	dy_url = "https://www.douyin.com/web/api/v2/aweme/iteminfo/?item_ids="
	// pc端
	pc_ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"
	// 移动端
	phone_ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
)

type IndexController struct{}

func (ctr *IndexController) Index(ctx *gin.Context) {
	var path string
	var keyWords string = ctx.Query("key_words")
	if keyWords == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "缺少关键词",
			"data": "",
		})
		return
	}
	url := modules.GetUrl(keyWords)
	if strings.Index(url, "douyin.com") != -1 {
		path = handleVideo.DouYin(url, phone_ua)
	} else if strings.Index(url, "b23.tv") != -1 {
		path = handleVideo.Bilibili(url, phone_ua)
	} else if strings.Index(url, "huoshan.com") != -1 {
		path = handleVideo.HuoShan(url, phone_ua)
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"code": http.StatusOK,
		"data": path,
		"msg":  "解析完成",
	})
}
