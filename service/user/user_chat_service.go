package service

import (
	"errors"
	"quick-talk/model"
	"time"
)

// CheckUserRole 检查权限
func CheckUserRole(id uint) (int, error) {
	var user model.User

	if err := model.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return model.USER_NO_AUTHORIZ, errors.New("找不到用户")
	}
	if user.MembershipDate != nil && user.MembershipDate.After(time.Now()) {
		return model.USER_VIP_NORMAL, nil
	} else if user.Balance > 0 {
		return model.USER_HAS_BALANCE, nil
	}
	return model.USER_NO_AUTHORIZ, nil
}

func UserPayBalance(id uint, num float64) error {
	var user model.User

	if err := model.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return errors.New("找不到用户")
	}
	if user.Balance < num {
		return errors.New("余额不足")
	}
	user.Balance = user.Balance - num
	if err := model.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
