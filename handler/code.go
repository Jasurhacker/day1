package handler

import (
	"test/test/controller"

	"github.com/gin-gonic/gin"
)
func Handler() {
	server := gin.Default()

	server.POST("send",controller.Post)
	server.GET("get",controller.Get)

	server.Run(":2010")
}