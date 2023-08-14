package handleVideo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"watermarkServer/modules"
)

const dy_url = "https://www.douyin.com/web/api/v2/aweme/iteminfo/?item_ids="

// 抖音
// 获取真实请求地址的path
func getDYRealityUrl(rUrl, ua string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", rUrl, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	fmt.Println(req.Header)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	return resp, nil
}

// 通过path获取到到id 请求dy_url
func DouYin(rUrl string, ua string) (string, error) {
	resp, err := getDYRealityUrl(rUrl, ua)
	if err != nil {
		return "", err
	}
	path := resp.Request.URL.Path
	re := regexp.MustCompile("[0-9]+")
	ids := re.FindAllString(path, -1)
	if len(ids) == 0 {
		return "", errors.New("无效链接！！！")
	}
	body, err := modules.HttpGet(dy_url+ids[0], ua)
	fmt.Println("获取视频", dy_url+ids[0], body, err)
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
	r, err := getDYRealityUrl(wmVideoUrl, ua)
	if err != nil {
		return "", err
	}
	vUrl, _ := url.QueryUnescape(r.Request.URL.String())
	return vUrl, nil
}
