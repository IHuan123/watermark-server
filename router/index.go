package router

import (
	"github.com/gin-gonic/gin"
	"watermarkServer/controllers"
)

var ctr = controllers.IndexController{}

func IndexRouter(r *gin.Engine) {
	r.GET("/", ctr.Index)
}
