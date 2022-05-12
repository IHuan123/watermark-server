package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"watermarkServer/router"
)

func main() {
	r := gin.Default()
	////// 记录错误日志到文件，同时输出到控制台
	fErr, _ := os.Create("log/gin_err.log")
	gin.DefaultErrorWriter = io.MultiWriter(fErr, os.Stdout)
	router.IndexRouter(r)
	if err := r.Run(":9000"); err != nil {
		panic(err)
	}
}
