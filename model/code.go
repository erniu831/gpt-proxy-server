package model

import (
	"gorm.io/gorm"
	"time"
)

type Code struct {
	gorm.Model
	Code       string `gorm:"uniqueIndex"`
	Type       int    `gorm:"default:0"`
	Value      string
	ExpireAt   time.Time
	Status     int `gorm:"default:0"`
	UsedUserId int
	UsedTime   time.Time
}

const (
	CODE_TYPE_DATE = iota
	CODE_TYPE_VALUE
)

const (
	CODE_STATUS_NEW = iota
	CODE_STATUS_USED
)
