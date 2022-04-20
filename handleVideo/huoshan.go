package handleVideo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"watermarkServer/modules"
)

const hs_url = "https://share.huoshan.com/api/item/info?item_id="

//获取真实请求地址的path
func getHSRealityUrl(url, ua string) url.Values {
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
	return resp.Request.URL.Query()
}

type HSRes struct {
	Data *struct {
		Item_info *struct {
			Cover string
			Url   string
		}
	}
}

//火山
func HuoShan(url, ua string) string {
	query := getHSRealityUrl(url, ua)
	fmt.Println(query)
	item_id := query["item_id"]
	fmt.Println(item_id)
	body := modules.HttpGet(hs_url+item_id[0], ua)
	var res HSRes
	json.Unmarshal(body, &res)
	return res.Data.Item_info.Url
}
