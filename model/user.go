package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// User 用户模型
type User struct {
	gorm.Model
	Username       string `gorm:"uniqueIndex,size:191"`
	PasswordDigest string
	Phone          string    `gorm:"uniqueIndex,size:191"`
	Email          string    `gorm:"uniqueIndex,size:191"`
	Membership     int       `gorm:"default:0"`
	MembershipDate time.Time `gorm:"default:null"`
	Balance        float64   `gorm:"default:0"`
	Status         string    `gorm:"default:null"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

const (
	USER_NO_AUTHORIZ = iota
	USER_HAS_BALANCE
	USER_VIP_NORMAL
	USER_VIP_PLUS
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
