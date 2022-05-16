package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"watermarkServer/router"
)

func main() {
	r := gin.Default()
	////// 记录错误日志到文件，同时输出到控制台
	fErr, _ := os.Create("log/gin_err.log")
	gin.DefaultErrorWriter = io.MultiWriter(fErr, os.Stdout)

	r.Static("/assets", "./assets")
	r.Use(static.Serve("/", static.LocalFile("template/dist", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("template/dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write(content)
			c.Writer.Flush()
		}
	})
	router.IndexRouter(r)
	if err := r.Run(":9000"); err != nil {
		panic(err)
	}
}
