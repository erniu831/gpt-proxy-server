package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"quick-talk/gpt"
	chatService "quick-talk/service/chat"
)

// UserRegister 用户注册接口
func ChatCompletion(c *gin.Context) {
	var service chatService.CompletionService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Println(fmt.Sprintf("%+v", service))
		gpt.ChatClient.Completion(c, service)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
