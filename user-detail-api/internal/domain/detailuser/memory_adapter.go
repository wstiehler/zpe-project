package detailuser

import (
	"errors"

	"gorm.io/gorm"
)

type MemorySqlAdapter struct {
}

func (MemorySqlAdapter) GetUserByEmail(db *gorm.DB, email string) (UserEntity, error) {
	var user UserEntity
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UserEntity{}, ErrUserNotFound
		}
		return UserEntity{}, err
	}
	return user, nil
}

func (MemorySqlAdapter) GetUserByID(db *gorm.DB, id string) (UserEntity, error) {
	var user UserEntity
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UserEntity{}, ErrUserNotFound
		}
		return UserEntity{}, err
	}
	return user, nil
}

func (MemorySqlAdapter) GetUserByName(db *gorm.DB, name string) (UserEntity, error) {
	var user UserEntity
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UserEntity{}, ErrUserNotFound
		}
		return UserEntity{}, err
	}
	return user, nil
}

func (MemorySqlAdapter) GetAllUsers(db *gorm.DB) ([]UserEntity, error) {
	var users []UserEntity
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
