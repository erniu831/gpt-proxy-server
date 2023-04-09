package api

import (
	"errors"
	"fmt"
	"quick-talk/gpt"
	"quick-talk/model"
	chatService "quick-talk/service/chat"
	userService "quick-talk/service/user"

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
	var userId uint
	cUser, exist := c.Get("user")
	if exist {
		user, ok := cUser.(*model.User)
		if ok {
			userId = user.ID
		}
	}
	role, err := userService.CheckUserRole(userId)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	if role == model.USER_NO_AUTHORIZ {
		c.JSON(200, ErrorResponse(errors.New("余额不足")))
		return
	}
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	var service chatService.CompletionService
	if err := c.ShouldBind(&service); err == nil {
		gpt.CompletionSSE(c, service)
		if err == nil {
			userService.UserPayBalance(userId, 1)
		}
	} else {
		c.SSEvent("complete", "error")
	}
}

// GET ChatCompletionSSE
func ChatCompletionForFreeSSE(c *gin.Context) {
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
