package code

import (
	"errors"
	"quick-talk/model"
	"quick-talk/serializer"
)

type CheckCodeService struct {
	Code   string `form:"code" json:"code"`
	UserId int    `form:"userId" json:"userId"`
}

func (service *CheckCodeService) CheckCode() serializer.Response {
	return CheckCode(*service)
}

func CheckCode(service CheckCodeService) serializer.Response {
	var dbCode model.Code
	if err := model.DB.Where("code = ?", service.Code).First(&dbCode).Error; err != nil {
		return serializer.Err(serializer.CodeYWError, "该兑换码不存在", errors.New("该兑换码不存在"))
	}

	return serializer.BuildCodeResponse(dbCode)
}
