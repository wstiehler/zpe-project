package user

import (
	"gorm.io/gorm"
)

type QueryAdapter interface {
	GetUserByEmail(db *gorm.DB, email string) (UserEntity, error)
	UpdateUser(db *gorm.DB, user *UserEntity, id string) (err error)
	GetUserById(db *gorm.DB, id string) (UserEntity, error)
	DeleteUser(db *gorm.DB, id string) error
}
