package handleVideo

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

//获取真实请求地址的path
func getBiLi(url, ua string) []byte {
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
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

//bilibili
func Bilibili(url, ua string) string {
	body := getBiLi(url, ua)
	reg := regexp.MustCompile(`readyVideoUrl:\s*'(.*)',`)
	regs := reg.FindStringSubmatch(string(body))
	if len(regs) != 2 {
		return ""
	}
	return regs[0]
}
