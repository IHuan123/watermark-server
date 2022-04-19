package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strings"
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
type Url struct {
	Url string
}

//获取真实请求地址的path
func getRealityUrl(url, ua string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	return resp.Request.URL.Path
}

// 抖音
//通过path获取到到id 请求dy_url
func douYin(url string) string {
	path := getRealityUrl(url, phone_ua)
	re := regexp.MustCompile("[0-9]+")
	ids := re.FindAllString(path, -1)
	if len(ids) == 0 {
		return ""
	}
	body := modules.HttpGet(dy_url+ids[0], phone_ua)

	type Video struct {
		Item_list []*struct {
			Video *struct {
				Play_addr *struct {
					Url_list []string
				}
			}
		}
	}
	var data Video
	json.Unmarshal([]byte(body), &data)
	fmt.Println(data.Item_list[0].Video.Play_addr.Url_list[0])
	wmVideoUrl := data.Item_list[0].Video.Play_addr.Url_list[0]
	wmVideoUrl = strings.Replace(wmVideoUrl, "playwm", "play", -1)
	return wmVideoUrl

}

func (ctr *IndexController) Index(ctx *gin.Context) {
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
	path := douYin(url)
	ctx.String(http.StatusOK, path)
}
