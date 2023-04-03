package api

import (
	"quick-talk/model"
	"quick-talk/serializer"
	codeService "quick-talk/service/code"

	"github.com/gin-gonic/gin"
)

// CodeGenerate 生成
func CodeGenerate(c *gin.Context) {
	var service codeService.GenertaeCodeService
	if err := c.ShouldBind(&service); err == nil {
		res, err := service.GenertaeCode()
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		c.JSON(200, serializer.BuildStringResponse(res))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CodeGenerate 兑换
func CodeRedeem(c *gin.Context) {
	var service codeService.RedeemCodeService
	if service.UserId == 0 {
		cUser, exist := c.Get("user")
		if exist {
			user, ok := cUser.(model.User)
			if ok {
				service.UserId = user.ID
			}
		}
	}
	if err := c.ShouldBind(&service); err == nil {
		err = service.RedeemCode()
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		c.JSON(200, serializer.BuildBoolResponse(true))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CodeCheck 查询
func CodeCheck(c *gin.Context) {
	var service codeService.CheckCodeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CheckCode()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
