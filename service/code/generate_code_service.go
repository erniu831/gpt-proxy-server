package code

import (
	"fmt"
	"math/rand"
	"quick-talk/model"
	"time"
)

type GenertaeCodeService struct {
	Type  int    `form:"type" json:"type"`
	Value string `form:"value" json:"value"`
}

func (service *GenertaeCodeService) GenertaeCode() (string, error) {
	return GenertaeCode(*service)
}

func GenertaeCode(service GenertaeCodeService) (string, error) {
	code := fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(10000))
	value := service.Value
	ct := service.Type
	expireInHours := 72
	expireAt := time.Now().Add(time.Duration(expireInHours) * time.Hour)

	dbCode := model.Code{Code: code, Type: ct, Value: value, ExpireAt: expireAt}
	if err := model.DB.Create(&dbCode).Error; err != nil {
		return "", err
	}
	return code, nil
}
