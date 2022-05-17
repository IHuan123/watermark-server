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
	var err error
	type Platform struct {
		Name     string `json:"name"`
		Platform string `json:"platform"`
	}
	var platformInfo = &Platform{}
	var keyWords string = ctx.Query("key_words")
	if keyWords == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "缺少关键词",
			"data": nil,
		})
		return
	}
	rUrl, err := modules.GetUrl(keyWords)
	if err != nil {
		base.Err(ctx, err.Error())
		return
	}
	if strings.Contains(rUrl, "douyin.com") { //抖音
		platformInfo.Name = "抖音"
		platformInfo.Platform = "douyin"
		path, err = handleVideo.DouYin(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "b23.tv") { //哔哩哔哩
		platformInfo.Name = "bilibili"
		platformInfo.Platform = "bilibili"
		path, err = handleVideo.BiliBili(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "huoshan.com") { //火山
		platformInfo.Name = "火山"
		platformInfo.Platform = "huoshan"
		path, err = handleVideo.HuoShan(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "kuaishou.com") { //快手
		platformInfo.Name = "快手"
		platformInfo.Platform = "kuaishou"
		path, err = handleVideo.KuaiShou(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "pipix.com") { // 皮皮虾
		platformInfo.Name = "皮皮虾"
		platformInfo.Platform = "pipix"
		path, err = handleVideo.PiPixia(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "weishi.qq.com") { //微视
		platformInfo.Name = "微视"
		platformInfo.Platform = "weishi"
		path, err = handleVideo.WeiShi(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "xiaochuankeji.cn") { //最右
		platformInfo.Name = "最右"
		platformInfo.Platform = "zuiyou"
		path, err = handleVideo.ZuiYou(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "m.eyepetizer.net") { //开眼
		platformInfo.Name = "开眼"
		platformInfo.Platform = "kaiyan"
		path, err = handleVideo.KaiYan(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "kg3.qq.com") { //全民k歌
		platformInfo.Name = "全民k歌"
		platformInfo.Platform = "kg"
		path, err = handleVideo.QuanMingKGe(rUrl, phone_ua)
	} else if strings.Contains(rUrl, "pipigx.com") { //皮皮搞笑
		platformInfo.Name = "皮皮搞笑"
		platformInfo.Platform = "pipigx"
		path, err = handleVideo.PiPiGX(rUrl, phone_ua)
	} else {
		platformInfo.Name = ""
		platformInfo.Platform = ""
		base.Err(ctx, "暂不支持该平台！！！")
		return
	}
	if err != nil {
		base.Err(ctx, err.Error())
		return
	}
	var res = &struct {
		PlatformInfo *Platform `json:"platformInfo"`
		Path         string    `json:"path"`
	}{
		PlatformInfo: platformInfo,
		Path:         path,
	}

	base.Success(ctx, res, "解析完成")
}
