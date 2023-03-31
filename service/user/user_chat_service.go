package service

import (
	"errors"
	"quick-talk/model"
	"time"
)

// Login 用户登录函数
func CheckUserRole(id int) (int, error) {
	var user model.User

	if err := model.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return model.USER_NO_AUTHORIZ, errors.New("找不到用户")
	}
	if user.MembershipDate.After(time.Now()) {
		return model.USER_VIP_NORMAL, nil
	} else if user.Balance > 0 {
		return model.USER_HAS_BALANCE, nil
	}
	return model.USER_NO_AUTHORIZ, nil
}

func UserPayBalance(id int, num float64) error {
	return nil
}
