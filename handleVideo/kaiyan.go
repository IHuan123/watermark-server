package handleVideo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"watermarkServer/modules"
)

var kaiYanUrl = "https://proxy.eyepetizer.net/v1/content/item/get_item_detail_v2"

//无法携带form参数 待解决
func requestKaiYan(rUrl, id, ua string) ([]byte, error) {
	client := &http.Client{}
	formValues := url.Values{}
	formValues.Set("resource_id", id)
	formValues.Set("resource_type", "pgc_video")
	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)
	request, err := http.NewRequest("POST", rUrl, formBytesReader)
	if err != nil {
		return nil, err
	}
	key := "token"
	dict := make(map[string]interface{})
	tokenNew, err := modules.GenerateToken(dict, key)
	if err != nil {
		return nil, err
	}
	//注意别忘了设置header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("user-agent", ua)
	request.Header.Set("x-thefair-appid", "xfpa44crf2p70lk8")
	request.Header.Set("x-thefair-auth", tokenNew)
	request.Header.Set("x-thefair-cid", "6ba20b639ea95388f36cfb468e111b1a")
	request.Header.Set("x-thefair-forward-host", "https://api.eyepetizer.net")
	request.Header.Set("x-thefair-ua", "EYEPETIZER_UNIAPP_H5/100000 (android;android;OS_VERSION_UNKNOWN;zh-Hans-CN;h5;1.0.0;cn-bj;SOURCE_UNKNOWN;6ba20b639ea95388f36cfb468e111b1a;2560*1080;NETWORK_UNKNOWN) native/1.0")
	do, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()
	body, err := ioutil.ReadAll(do.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func KaiYan(sUrl, ua string) (string, error) {
	sUri, err := url.Parse(sUrl)
	if err != nil {
		return "", err
	}
	query := sUri.Query()
	ids := query["video_id"]
	fmt.Println(ids)
	if len(ids) == 1 {
		body, err := requestKaiYan(kaiYanUrl, ids[0], ua)
		if err != nil {
			return "", err
		}
		var res struct {
			Code   int `json:"code"`
			Result struct {
				Video struct {
					PlayUrl string `json:"play_url"`
				} `json:"video"`
			} `json:"result"`
		}
		err = json.Unmarshal(body, &res)
		if err != nil {
			return "", err
		}
		if res.Code != 0 {
			return "", errors.New("解析出现问题！")
		}
		return res.Result.Video.PlayUrl, nil
	} else {
		return "", errors.New("无效的地址")
	}

}
