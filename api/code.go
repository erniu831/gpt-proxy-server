package api

import (
	codeService "quick-talk/service/code"

	"github.com/gin-gonic/gin"
)

// CodeGenerate 兑换
func CodeGenerate(c *gin.Context) {
	var service codeService.GenertaeCodeService
	if err := c.ShouldBind(&service); err == nil {
		res, err := service.GenertaeCode()
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CodeGenerate 兑换
func CodeRedeem(c *gin.Context) {
	var service codeService.RedeemCodeService
	if err := c.ShouldBind(&service); err == nil {
		err = service.RedeemCode()
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		c.JSON(200, true)
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
