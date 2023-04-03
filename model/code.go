package model

import (
	"time"

	"gorm.io/gorm"
)

type Code struct {
	gorm.Model
	Code       string `gorm:"uniqueIndex,size:191"`
	Type       int    `gorm:"default:0"`
	Value      string
	ExpireAt   time.Time
	Status     int `gorm:"default:0"`
	UsedUserId uint
	UsedTime   time.Time `gorm:"default:null"`
}

const (
	CODE_TYPE_DATE = iota
	CODE_TYPE_VALUE
)

const (
	CODE_STATUS_NEW = iota
	CODE_STATUS_USED
)
