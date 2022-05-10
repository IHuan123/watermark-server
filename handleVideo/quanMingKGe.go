package handleVideo

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

func requestKG3(rUrl, ua string) ([]byte, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", rUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("user-agent", ua)
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

func QuanMingKGe(sUrl, ua string) (string, error) {

	s := regexp.MustCompile(`s=(.*?)&`).FindStringSubmatch(sUrl)
	if len(s) != 2 {
		return "", errors.New("无效的地址")
	}
	body, err := requestKG3("https://kg.qq.com/node/play?s="+s[1], ua)
	if err != nil {
		return "", err
	}
	regs := regexp.MustCompile(`playurl_video":"(.*?)","poi_id`).FindStringSubmatch(string(body))
	if len(regs) != 2 {
		return "", errors.New("无效的地址")
	}
	return regs[1], nil
}
