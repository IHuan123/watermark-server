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
	rUrl := modules.GetUrl(keyWords)
	var err error
	if strings.Contains(rUrl, "douyin.com") { //抖音
		path, err = handleVideo.DouYin(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "b23.tv") { //哔哩哔哩
		path, err = handleVideo.BiliBili(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "huoshan.com") { //火山
		path, err = handleVideo.HuoShan(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "kuaishou.com") { //快手
		path, err = handleVideo.KuaiShou(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "pipix.com") { // 皮皮虾
		path, err = handleVideo.PiPixia(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "weishi.qq.com") { //微视
		path, err = handleVideo.WeiShi(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "xiaochuankeji.cn") { //最右
		path, err = handleVideo.ZuiYou(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "m.eyepetizer.net") { //开眼
		path, err = handleVideo.KaiYan(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "kg3.qq.com") { //全民k歌
		path, err = handleVideo.QuanMingKGe(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "pipigx.com") { //皮皮搞笑
		path, err = handleVideo.PiPiGX(rUrl, phone_ua)
	} else {
		base.Err(ctx, "暂不支持该平台！！！")
	}
	if err != nil {
		base.Err(ctx, err.Error())
		return
	}

	base.Success(ctx, path, "解析完成")
}
