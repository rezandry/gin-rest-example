package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rezandry/rest-api-gin/controller"
)

func main() {
	port := "5001"
	r := gin.Default()
	r.POST("/login", controller.Login)
	r.Use(controller.CheckHeader)
	r.GET("/:id", controller.GetHome)
	r.Run(":" + port)
}
