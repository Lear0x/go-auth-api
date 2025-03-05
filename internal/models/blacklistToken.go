package models

import "gorm.io/gorm"

type BlacklistedToken struct {
	gorm.Model
	Token string `gorm:"unique"`
}
