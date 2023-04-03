package serializer

import (
	"quick-talk/model"
	"time"
)

// User 用户序列化器
type Code struct {
	Code       string    `json:"code"`
	Type       int       `json:"type"`
	Value      string    `json:"value"`
	ExpireAt   time.Time `json:"expireAt"`
	Status     int       `json:"status"`
	UsedUserId uint      `json:"usedUserId"`
	UsedTime   time.Time `json:"usedTime"`
}

// BuildUser 序列化用户
func BuildCode(code model.Code) Code {
	return Code{
		Code:       code.Code,
		Type:       code.Type,
		Value:      code.Value,
		ExpireAt:   code.ExpireAt,
		Status:     code.Status,
		UsedUserId: code.UsedUserId,
		UsedTime:   code.UsedTime,
	}
}

// BuildUserResponse 序列化用户响应
func BuildCodeResponse(user model.Code) Response {
	return Response{
		Data: BuildCode(user),
	}
}
