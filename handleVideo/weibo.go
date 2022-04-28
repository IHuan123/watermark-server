package handleVideo

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

func getWeiBoBody(rurl, id, ua string) (string, error) {
	postData := `{"Component_Play_Playinfo":{"oid":"` + id + `"}}`
	formValues := url.Values{}
	formValues.Add("data", postData)
	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)
	req, err := http.NewRequest("POST", rurl, formBytesReader)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "login_sid_t=6b652c77c1a4bc50cb9d06b24923210d; cross_origin_proto=SSL; WBStorage=2ceabba76d81138d|undefined; _s_tentry=passport.weibo.com; Apache=7330066378690.048.1625663522444; SINAGLOBAL=7330066378690.048.1625663522444; ULV=1625663522450:1:1:1:7330066378690.048.1625663522444:; TC-V-WEIBO-G0=35846f552801987f8c1e8f7cec0e2230; SUB=_2AkMXuScYf8NxqwJRmf8RzmnhaoxwzwDEieKh5dbDJRMxHRl-yT9jqhALtRB6PDkJ9w8OaqJAbsgjdEWtIcilcZxHG7rw; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9W5Qx3Mf.RCfFAKC3smW0px0; XSRF-TOKEN=JQSK02Ijtm4Fri-YIRu0-vNj")
	req.Header.Set("referer", "https://weibo.com/tv/show/"+id)
	//req.Header.Set("accept-encoding", "gzip, deflate")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("origin", "https://weibo.com")
	req.Header.Set("page-referer", "/tv/show/"+id)
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"`)
	req.Header.Set("sec-ch-ua-platform", "macOS")
	req.Header.Set("x-xsrf-token", "fQIve1IUgGs47DN9qnee4CPd")
	client := &http.Client{}
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
	//parse, err := url.Parse(vUrl)
	//if err != nil {
	//	return "", err
	//}
	if len(ids) != 2 {
		return "", errors.New("无效地址")
	}
	//url.QueryEscape()
	rUrl := "https://weibo.com/api/component?page=/tv/show/" + ids[0]
	body, err := getWeiBoBody(rUrl, ids[0], ua)
	if err != nil {
		return "", err
	}
	return body, nil
}
