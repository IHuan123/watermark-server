package modules

import (
	"mvdan.cc/xurls/v2"
)

var types [1]string = [1]string{"douyin"}

func urlType(url string) string {

	return ""
}

// 提取文字中的url
func GetUrl(url string) string {
	rxRelaxed := xurls.Strict()
	src := rxRelaxed.FindAllString(url, -1)
	return src[0]
}
