package handleVideo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//获取真实请求地址的path
func getKSRealityUrl(url, ua string) []byte {
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
func KuaiShou(url, ua string) string {
	body := getKSRealityUrl(url, ua)
	fmt.Println(string(body))
	return string(body)
}
