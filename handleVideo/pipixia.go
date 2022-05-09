package handleVideo

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"watermarkServer/modules"
)

func getPiPiXia(rUrl string, ua string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", rUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	fmt.Println(resp.Request.URL.RequestURI(), resp.Request.URL.Path)
	path := resp.Request.URL.Path
	return path, nil
}

func PiPixia(rUrl string, ua string) (string, error) {
	path, err := getPiPiXia(rUrl, ua)
	if err != nil {
		return "", err
	}
	reg := regexp.MustCompile("[0-9]+")
	ids := reg.FindAllString(path, -1)
	if len(ids) == 0 {
		return "", errors.New("无效地址")
	}
	infoUrl := "https://is.snssdk.com/bds/cell/detail/?cell_type=1&aid=1319&app_name=super&cell_id=" + ids[0]
	body, err := modules.HttpGet(infoUrl, ua)
	if err != nil {
		return "", err
	}
	reg2 := regexp.MustCompile("origin_video_download.*?url_list.*?url.*?:\"(.*?)\"") // re.findall("origin_video_download.*?url_list.*?url.*?:\"(.*?)\"",r.text)
	videoUrl := reg2.FindStringSubmatch(string(body))[1]
	return videoUrl, nil
}
