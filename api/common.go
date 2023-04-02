package api

import (
	"errors"
	"quick-talk/model"
	"quick-talk/serializer"

	"github.com/gin-gonic/gin"
)

// UserLogin 用户登录接口
func GetSession(c *gin.Context) {
	user, exists := c.Get("user")
	if exists {
		userModel, _ := user.(*model.User)
		// _ = json.Unmarshal([]byte(user), &userModel)
		c.JSON(200, serializer.BuildUserResponse(*userModel))
	} else {
		c.JSON(200, ErrorResponse(errors.New("未登录")))
	}
}
