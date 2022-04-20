package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"watermarkServer/modules"
	"watermarkServer/router"
)

func getUrl(keyWords string) []byte {
	url := modules.GetUrl(keyWords)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
func main() {
	fmt.Println("水印")
	r := gin.Default()
	////// 记录错误日志到文件，同时输出到控制台
	fErr, _ := os.Create("log/gin_err.log")
	gin.DefaultErrorWriter = io.MultiWriter(fErr, os.Stdout)
	router.IndexRouter(r)
	if err := r.Run(":9000"); err != nil {
		panic(err)
	}
}
