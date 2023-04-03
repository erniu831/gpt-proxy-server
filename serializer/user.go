package serializer

import (
	"quick-talk/model"
	"time"
)

// User 用户序列化器
type User struct {
	ID             uint      `json:"id"`
	Username       string    `json:"username"`
	Status         string    `json:"status"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	Membership     int       `json:"membership"`
	MembershipDate time.Time `json:"membershipDate"`
	Balance        float64   `json:"balance"`
	CreatedAt      int64     `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:             user.ID,
		Username:       user.Username,
		Status:         user.Status,
		CreatedAt:      user.CreatedAt.Unix(),
		Phone:          user.Phone,
		Email:          user.Email,
		Membership:     user.Membership,
		MembershipDate: *user.MembershipDate,
		Balance:        user.Balance,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
