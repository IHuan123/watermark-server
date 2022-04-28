package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"watermarkServer/handleVideo"
	"watermarkServer/modules"
)

const (
	// pc端
	pc_ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"
	//mac chrome
	mac_ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"
	// 移动端
	phone_ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
)

type IndexController struct{}

var base = &modules.BaseController{}

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
	var err error
	if strings.Contains(url, "douyin.com") {
		path, err = handleVideo.DouYin(url, phone_ua)
	} else if strings.Contains(url, "b23.tv") {
		path, err = handleVideo.BiliBili(url, phone_ua)
	} else if strings.Contains(url, "huoshan.com") {
		path, err = handleVideo.HuoShan(url, phone_ua)
	} else if strings.Contains(url, "kuaishou.com") {
		path, err = handleVideo.KuaiShou(url, phone_ua)
	} else if strings.Contains(url, "pipix.com") {
		path, err = handleVideo.PiPixia(url, phone_ua)
	} else if strings.Contains(url, "weibo.com") {
		path, err = handleVideo.WeiBo(url, pc_ua)
	} else if strings.Contains(url, "weishi.qq.com") {
		path, err = handleVideo.WeiShi(url, phone_ua)
	} else if strings.Contains(url, "xiaochuankeji.cn") {
		path, err = handleVideo.ZuiYou(url, phone_ua)
	}
	if err != nil {
		base.Err(ctx, err.Error())
		return
	}

	base.Success(ctx, path, "解析完成")
}
