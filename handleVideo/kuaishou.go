package handleVideo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getKSRealityUrl(url, ua string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Request.URL)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//bilibili
func KuaiShou(url, ua string) (string, error) {
	body, err := getKSRealityUrl(url, ua)
	if err != nil {
		return "", err
	}
	//解析获取的body
	regs := regexp.MustCompile(`srcNoMark":"(.*?)"`).FindStringSubmatch(string(body))
	if len(regs) != 2 {
		return "", errors.New("无效地址")
	}
	return regs[1], nil
}
