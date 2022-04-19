package modules

import (
	"mvdan.cc/xurls/v2"
)

// 提取文字中的url
func GetUrl(url string) string {
	rxRelaxed := xurls.Strict()
	src := rxRelaxed.FindAllString(url, -1)
	return src[0]
}
