package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"watermarkServer/modules"
)

type IndexController struct{}
type Url struct {
	Url string
}

func (ctr *IndexController) Index(ctx *gin.Context) {
	var keyWords string = ctx.Query("key_words")
	if keyWords == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "缺少关键词",
			"data": "",
		})
		return
	}
	url := modules.GetUrl(keyWords)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// 写入文件
	ioutil.WriteFile("site.txt", body, 0644)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": string(body),
		"msg":  "success",
	})
	//ctx.JSON(http.StatusOK, string(body))
}
