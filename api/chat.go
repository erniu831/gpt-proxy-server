package api

import (
	"fmt"
	"quick-talk/gpt"
	chatService "quick-talk/service/chat"

	"github.com/gin-gonic/gin"

	gogpt "github.com/sashabaranov/go-openai"
)

// UserRegister 用户注册接口
func ChatCompletion(c *gin.Context) {
	var service chatService.CompletionService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Println(fmt.Sprintf("%+v", service))
		res, err := gpt.Completion(c, service)
		if err != nil {
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ChatProcess(c *gin.Context) {
	var service chatService.ChatProcessService
	if err := c.ShouldBind(&service); err == nil {
		fmt.Println(fmt.Sprintf("%+v", service))
		res, err := gpt.Completion(c, chatService.CompletionService{
			Messages: []gogpt.ChatCompletionMessage{
				{
					Role:    "user",
					Content: service.Prompt,
				},
			},
		})
		if err != nil {
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
