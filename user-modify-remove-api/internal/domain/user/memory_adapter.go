package user

import (
	"gorm.io/gorm"
)

type MemorySqlAdapter struct {
}

func (MemorySqlAdapter) GetUserByEmail(db *gorm.DB, email string) (UserEntity, error) {
	var user UserEntity
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return UserEntity{}, err
	}
	return user, nil
}
