package service

import (
	"gorm.io/gorm"
)

type UserService interface {
	GetUser(db *gorm.DB)
	Register(db *gorm.DB) int
}
