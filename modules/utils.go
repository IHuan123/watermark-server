package modules

import (
	"errors"
	"fmt"
	"mvdan.cc/xurls/v2"
)

// 提取文字中的url
func GetUrl(url string) (string, error) {
	rxRelaxed := xurls.Strict()
	src := rxRelaxed.FindAllString(url, -1)
	fmt.Println(src)
	if len(src) == 0 {
		return "", errors.New("无效地址")
	}
	return src[0], nil
}
