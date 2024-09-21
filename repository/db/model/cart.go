package model

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint `gorm:"not null"`
	Num       uint
	MaxNum    uint
	Check     bool
}
