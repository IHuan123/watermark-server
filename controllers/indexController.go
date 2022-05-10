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
	if strings.Contains(url, "douyin.com") { //抖音
		path, err = handleVideo.DouYin(url, phone_ua)
	} else if strings.Contains(url, "b23.tv") { //哔哩哔哩
		path, err = handleVideo.BiliBili(url, phone_ua)
	} else if strings.Contains(url, "huoshan.com") { //火山
		path, err = handleVideo.HuoShan(url, phone_ua)
	} else if strings.Contains(url, "kuaishou.com") { //快手
		path, err = handleVideo.KuaiShou(url, phone_ua)
	} else if strings.Contains(url, "pipix.com") { // 皮皮虾
		path, err = handleVideo.PiPixia(url, phone_ua)
	} else if strings.Contains(url, "weishi.qq.com") { //微视
		path, err = handleVideo.WeiShi(url, phone_ua)
	} else if strings.Contains(url, "xiaochuankeji.cn") { //最右
		path, err = handleVideo.ZuiYou(url, phone_ua)
	} else if strings.Contains(url, "m.eyepetizer.net") { //开眼
		path, err = handleVideo.KaiYan(url, phone_ua)
	} else if strings.Contains(url, "kg3.qq.com") { //全民k歌
		path, err = handleVideo.QuanMingKGe(url, phone_ua)
	} else if strings.Contains(url, "pipigx.com") { //皮皮搞笑
		path, err = handleVideo.PiPiGX(url, phone_ua)
	} else {
		base.Err(ctx, "暂不支持该平台！！！")
	}
	if err != nil {
		base.Err(ctx, err.Error())
		return
	}

	base.Success(ctx, path, "解析完成")
}
