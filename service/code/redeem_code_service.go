package code

import (
	"errors"
	"quick-talk/model"
	"strconv"
	"time"
)

type RedeemCodeService struct {
	Code   string `form:"code" json:"code" binding: "required"`
	UserId uint   `form:"userId" json:"userId"`
}

func (service *RedeemCodeService) RedeemCode() error {
	return RedeemCode(*service)
}

func RedeemCode(service RedeemCodeService) error {
	var dbCode model.Code
	if err := model.DB.Where("code = ?", service.Code).First(&dbCode).Error; err != nil {
		return errors.New("该兑换码不存在")
	}
	if dbCode.Status == model.CODE_STATUS_USED {
		return errors.New("该兑换码已被核销")
	}

	var user model.User
	userID := service.UserId
	if err := model.DB.First(&user, userID).Error; err != nil {
		return errors.New("用户不存在")
	}

	if dbCode.Type == model.CODE_TYPE_DATE {
		// Membership code
		value, err := strconv.Atoi(dbCode.Value)
		if err != nil {
			return errors.New("Invalid code value")
		}
		user.Membership = 1
		if user.MembershipDate.After(time.Now()) {
			user.MembershipDate = user.MembershipDate.Add(time.Duration(value) * time.Hour)
		} else {
			user.MembershipDate = time.Now().Add(time.Duration(value) * time.Hour)
		}
	} else if dbCode.Type == model.CODE_TYPE_VALUE {
		// Balance code
		balance, err := strconv.ParseFloat(dbCode.Value, 64)
		if err != nil {
			return errors.New("Invalid code value")
		}
		user.Balance += balance
	} else {
		return errors.New("Invalid code value")
	}

	if err := model.DB.Save(&user).Error; err != nil {
		return errors.New("更新用户信息失败")
	}
	dbCode.Status = model.CODE_STATUS_USED
	dbCode.UsedUserId = userID
	dbCode.UsedTime = time.Now()
	if err := model.DB.Save(&dbCode).Error; err != nil {
		return nil
	}
	return nil
}
