package api

import (
	"fmt"
	"quick-talk/gpt"
	chatService "quick-talk/service/chat"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func ChatCompletion(c *gin.Context) {
	var service chatService.CompletionService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Println(fmt.Sprintf("%+v", service))
		res, err := gpt.Completion(c, service)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, ErrorResponse(err))
		} else {
			fmt.Println(res)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GET ChatCompletionSSE
func ChatCompletionSSE(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	var service chatService.CompletionService
	if err := c.ShouldBind(&service); err == nil {
		gpt.CompletionSSE(c, service)
	} else {
		c.SSEvent("complete", "error")
	}
}
