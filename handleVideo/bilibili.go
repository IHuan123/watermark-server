package handleVideo

import (
	"fmt"
	"log"
	"net/http"
)

//获取真实请求地址的path
func getBiLiRealityUrl(url, ua string) string {
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
	fmt.Println(resp.Request.URL)
	return resp.Request.URL.Path
}

//bilibili
func Bilibili(url, ua string) string {
	path := getBiLiRealityUrl(url, ua)
	fmt.Println(path)
	return "hello"
}
