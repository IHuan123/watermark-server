package modules

import (
	"reflect"
	"testing"
)

func TestGetUrl(t *testing.T) {
	t_url := "2.51 DUy:/ 隔着屏幕都心疼！新消防员下队后坚持加练，只为在救援现场与生命赛跑……（供稿：浙江消防）%浙江dou知道 %消防 %致敬   https://v.douyin.com/NK8t7Ra/ 复制此链接，打开Dou音搜索，直接观看视频！"
	got, _ := GetUrl(t_url)
	want := "https://v.douyin.com/NK8t7Ra/"
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}
