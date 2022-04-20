package handleVideo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"watermarkServer/modules"
)

const dy_url = "https://www.douyin.com/web/api/v2/aweme/iteminfo/?item_ids="

// 抖音
//获取真实请求地址的path
func getDYRealityUrl(url, ua string) string {
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

//通过path获取到到id 请求dy_url
func DouYin(url string, ua string) string {
	path := getDYRealityUrl(url, ua)
	re := regexp.MustCompile("[0-9]+")
	ids := re.FindAllString(path, -1)
	if len(ids) == 0 {
		return ""
	}
	body := modules.HttpGet(dy_url+ids[0], ua)

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
