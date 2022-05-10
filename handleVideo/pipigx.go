package handleVideo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const piPiGXUrl = "https://h5.pipigx.com/ppapi/share/fetch_content"

func requestPiPiGX(rUrl string, ids []string, ua string) ([]byte, error) {
	data := make(map[string]interface{})
	data["pid"], _ = strconv.Atoi(ids[0])
	data["mid"], _ = strconv.Atoi(ids[1])
	data["type"] = "post"
	bytesData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", rUrl, bytes.NewBuffer(bytesData))
	req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
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

func PiPiGX(rUrl, ua string) (string, error) {
	parse, err := url.Parse(rUrl)
	if err != nil {
		return "", err
	}
	query := parse.Query()
	if len(query["mid"]) == 0 {
		return "", errors.New("无效地址")
	}
	mid := query["mid"][0]
	paths := strings.Split(parse.Path, "/")
	pid := paths[len(paths)-1]
	body, err := requestPiPiGX(piPiGXUrl, []string{pid, mid}, ua)
	if err != nil {
		return "", err
	}
	var res *struct {
		Data *struct {
			Post *struct {
				Videos map[string]interface{} `json:"videos"`
			} `json:"post"`
		} `json:"data"`
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}
	videos := res.Data.Post.Videos
	var info interface{}
	for _, videoInfo := range videos {
		info = videoInfo
	}
	var videoUrl string
	if info, ok := info.(map[string]interface{}); ok {
		videoUrl = info["url"].(string)
	} else {
		return "", errors.New("解析失败")
	}
	return videoUrl, nil
}
