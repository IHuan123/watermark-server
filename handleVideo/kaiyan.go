package handleVideo

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// https://proxy.eyepetizer.net/v1/content/item/get_item_detail_v2
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
	//注意别忘了设置header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("user-agent", ua)
	do, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	fmt.Println(do.Request.Form)
	defer do.Body.Close()
	body, err := ioutil.ReadAll(do.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func KaiYan(sUrl, ua string) (string, error) {
	fmt.Println(sUrl)
	sUri, err := url.Parse(sUrl)
	if err != nil {
		return "", err
	}
	query := sUri.Query()
	ids := query["video_id"]
	if len(ids) == 1 {
		body, err := requestKaiYan(kaiYanUrl, ids[0], ua)
		if err != nil {
			return "", err
		}
		return string(body), nil
	} else {
		return "", errors.New("无效的地址")
	}

}
