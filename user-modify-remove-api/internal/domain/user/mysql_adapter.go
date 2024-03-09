package user

import (
	"gorm.io/gorm"
)

type MysqlAdapter struct {
}

func (MysqlAdapter) GetUserByEmail(db *gorm.DB, email string) (UserEntity, error) {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user UserEntity
	if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
		tx.Rollback()
		return UserEntity{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return UserEntity{}, err
	}

	return user, nil
}

func (MysqlAdapter) UpdateUser(db *gorm.DB, user *UserEntity, userId string) error {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&UserEntity{}).Where("id = ?", userId).Updates(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (MysqlAdapter) GetUserById(db *gorm.DB, id string) (UserEntity, error) {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user UserEntity
	if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
		tx.Rollback()
		return UserEntity{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return UserEntity{}, err
	}

	return user, nil
}

func (MysqlAdapter) DeleteUser(db *gorm.DB, id string) error {
	err := db.Where("id = ?", id).Delete(&UserEntity{}).Error

	if err != nil {
		return err
	}
	return nil
}
