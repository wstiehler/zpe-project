package createuser

import (
	"gorm.io/gorm"
)

type MemorySqlAdapter struct {
}

func (MemorySqlAdapter) CreateUser(db *gorm.DB, user *UserEntity) (UserEntity, error) {
	if err := db.Create(user).Error; err != nil {
		return UserEntity{}, err
	}
	return *user, nil
}

func (MemorySqlAdapter) GetUserByEmail(db *gorm.DB, email string) (UserEntity, error) {
	var user UserEntity
	db.Where("email = ?", email).First(&user)
	return user, nil
}
