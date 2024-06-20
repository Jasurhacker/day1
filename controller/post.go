package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Messages struct{
	Message string
	Time string
	Status string
}

var slice []Messages

func Post(c *gin.Context) {
	// message := c.Request.FormValue("message")
	var messagesshablon Messages
	c.ShouldBindJSON(&messagesshablon)
	fmt.Printf("messagesshablon: %v\n", messagesshablon)

	slice = append(slice, Messages{
		Message: messagesshablon.Message,
		Status: "notread",
		Time: time.Now().String(),
	})

	c.JSON(200,"sucsess")
}