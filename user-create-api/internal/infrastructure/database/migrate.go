package config

import (
	"github.com/wstiehler/zpecreateuser-api/internal/domain/createuser"
	"gorm.io/gorm"
)

func AutoMigrateTables(db *gorm.DB) error {
	err := db.Table("users").AutoMigrate(&createuser.UserEntity{})
	if err != nil {
		return err
	}
	return nil
}
