package config

import (
	"github.com/wstiehler/role-api/internal/domain/role"
	"gorm.io/gorm"
)

func AutoMigrateTables(db *gorm.DB) error {
	err := db.Table("roles").AutoMigrate(&role.RoleEntity{})
	if err != nil {
		return err
	}
	err = db.Table("permissions").AutoMigrate(&role.PermissionEntity{})
	if err != nil {
		return err
	}
	return nil
}
