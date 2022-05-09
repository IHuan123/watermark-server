package handleVideo

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"watermarkServer/modules"
)

const hs_url = "https://share.huoshan.com/api/item/info?item_id="

//获取真实请求地址的path
func getHSRealityUrl(rUrl, ua string) (url.Values, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", rUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	return resp.Request.URL.Query(), nil
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
func HuoShan(rUrl, ua string) (string, error) {
	query, err := getHSRealityUrl(rUrl, ua)
	if err != nil {
		return "", err
	}
	item_id := query["item_id"]
	body, err := modules.HttpGet(hs_url+item_id[0], ua)
	if err != nil {
		return "", err
	}
	var res HSRes
	json.Unmarshal(body, &res)
	return res.Data.Item_info.Url, nil
}
