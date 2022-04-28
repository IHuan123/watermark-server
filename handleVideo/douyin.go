package handleVideo

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"
	"watermarkServer/modules"
)

const dy_url = "https://www.douyin.com/web/api/v2/aweme/iteminfo/?item_ids="

// 抖音
//获取真实请求地址的path
func getDYRealityUrl(url, ua string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	return resp.Request.URL.Path, nil
}

//通过path获取到到id 请求dy_url
func DouYin(url string, ua string) (string, error) {
	path, err := getDYRealityUrl(url, ua)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile("[0-9]+")
	ids := re.FindAllString(path, -1)
	if len(ids) == 0 {
		return "", errors.New("无效连接")
	}
	body, err := modules.HttpGet(dy_url+ids[0], ua)
	if err != nil {
		return "", err
	}
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
	wmVideoUrl := data.Item_list[0].Video.Play_addr.Url_list[0]
	wmVideoUrl = strings.Replace(wmVideoUrl, "playwm", "play", -1)
	return wmVideoUrl, nil
}
