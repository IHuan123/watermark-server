package handleVideo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const zuiYouUrl = "https://share.xiaochuankeji.cn/planck/share/post/detail"

func requestZuiYou(rUrl, id, ua string) ([]byte, error) {
	//var data map[string]interface{}
	//data["pid"] = id
	data := fmt.Sprintf("{\"pid\":%s}", id)
	//responseData, err := json.Marshal(data)
	//if err != nil {
	//	return nil, err
	//}
	req, err := http.NewRequest("POST", rUrl, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func ZuiYou(rUrl, ua string) (string, error) {
	parse, err := url.Parse(rUrl)
	if err != nil {
		return "", err
	}
	query := parse.Query()
	fmt.Println(query)
	if len(query["pid"]) == 0 {
		return "", errors.New("无效地址")
	}
	pid := query["pid"][0]
	body, err := requestZuiYou(zuiYouUrl, pid, ua)
	if err != nil {
		return "", err
	}
	type Video struct {
		Data *struct {
			Post *struct {
				Videos *struct {
					Urlext string
				}
			}
		}
	}
	var res Video
	//"urlsrc":"http://video.izuiyou.com/zyvd/264/5f/0e/2b59-b687-11ec-9b89-00163e0e67b8",
	//"urlext":"http://video.izuiyou.com/zyvd/264/5f/0e/2b59-b687-11ec-9b89-00163e0e67b8",
	//"urlwm":"http://video.izuiyou.com/zyvd/264/5f/0e/2b59-b687-11ec-9b89-00163e0e67b8",
	json.Unmarshal(body, &res)
	fmt.Println(res.Data.Post.Videos.Urlext)
	return string(body), nil
}
