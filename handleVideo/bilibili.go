package handleVideo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//获取真实请求地址的path
func getBiLi(url, ua string) ([]byte, error) {
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
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

//bilibili 画质低
func BiliBili(url, ua string) (string, error) {
	body, err := getBiLi(url, ua)
	if err != nil {
		return "", err
	}
	reg := regexp.MustCompile(`readyVideoUrl:\s*'(.*)',`)
	regs := reg.FindStringSubmatch(string(body))
	fmt.Println(len(regs))
	if len(regs) != 2 {
		return "", errors.New("无效地址")
	}
	return regs[1], nil
}
