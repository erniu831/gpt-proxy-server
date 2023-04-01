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
			c.JSON(200, res)
		} else {
			fmt.Println(err)
			c.JSON(200, ErrorResponse(err))
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
