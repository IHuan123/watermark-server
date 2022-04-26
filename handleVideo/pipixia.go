package handleVideo

import (
	"fmt"
	"net/http"
	"regexp"
)

func getPiPiXia(url string, ua string) string {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error()
	}
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	fmt.Println(resp.Request.URL.RequestURI(), resp.Request.URL.Path)
	path := resp.Request.URL.Path
	return path
}

func PiPixia(url string, ua string) string {
	path := getPiPiXia(url, ua)
	reg := regexp.MustCompile("[0-9]+")
	itemId := reg.FindAllString(path, -1)
	return itemId[0]
}
