package detailuser

import (
	"gorm.io/gorm"
)

type QueryAdapter interface {
	GetUserByEmail(db *gorm.DB, email string) (UserEntity, error)
	GetUserByID(db *gorm.DB, id string) (UserEntity, error)
	GetUserByName(db *gorm.DB, name string) (UserEntity, error)
	GetAllUsers(db *gorm.DB) ([]UserEntity, error)
}
