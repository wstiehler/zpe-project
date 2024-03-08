package createuser

import (
	"gorm.io/gorm"
)

type QueryAdapter interface {
	CreateUser(db *gorm.DB, product *UserEntity) (UserEntity, error)
	GetUserByEmail(db *gorm.DB, email string) (UserEntity, error)
}
