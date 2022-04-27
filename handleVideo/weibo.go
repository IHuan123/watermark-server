package handleVideo

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

func getWeiBoBody(rurl, id, ua string) (string, error) {
	client := &http.Client{}
	postData := `{"oid":"` + id + `"}`
	formValues := url.Values{}
	formValues.Set("Component_Play_Playinfo", postData)
	fmt.Println(formValues.Get("Component_Play_Playinfo"))
	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)
	req, err := http.NewRequest("POST", rurl, formBytesReader)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", ua)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Cookie", "_s_tentry=-; Apache=5794002554431.063.1616336111827; SINAGLOBAL=5794002554431.063.1616336111827; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WFzvSpxg3QF9rqPyAo0uG3y; UOR=,,www.baidu.com; ULV=1651068799971:1:1:1:5794002554431.063.1616336111827:; YF-V-WEIBO-G0=b09171a17b2b5a470c42e2f713edace0; SUB=_2AkMVNdyGf8NxqwJRmP0SyWPmZYh_ywnEieKjaS1dJRMxHRl-yj92qlMvtRB6PrXyaYJmixnw1_9lq6k8IDYU0VWcHbt6")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
func WeiBo(vUrl string, ua string) (string, error) {
	reg := regexp.MustCompile(`(\d+:{1}\d+)`)
	ids := reg.FindStringSubmatch(vUrl)
	parse, err := url.Parse(vUrl)
	if err != nil {
		return "", err
	}
	fmt.Println(parse.Path)
	if len(ids) != 2 {
		return "", errors.New("无效地址")
	}
	rUrl := "https://h5.video.weibo.com/api/component?page=" + parse.Path + "/" + ids[0]

	fmt.Println(rUrl)
	body, err := getWeiBoBody(rUrl, ids[0], ua)
	if err != nil {
		return "", err
	}
	fmt.Println(body)
	return "", nil
}
