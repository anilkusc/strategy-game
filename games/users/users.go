package users

import (
	"gorm.io/gorm"
)

type IUser interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]User, error)
}

type User struct {
	gorm.Model
	Username string
	Password string
	Role     string
}
