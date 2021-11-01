package routers

import (
	"source-server/controllers"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	route := gin.Default()
	route.Static("/static", "./static")

	route.GET("/js/*filepath", controllers.GetParams)
	route.POST("/upload", controllers.Upload)
	route.Run(":8083")
	return route
}
