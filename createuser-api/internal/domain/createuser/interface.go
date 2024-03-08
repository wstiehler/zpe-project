package createuser

import (
	"gorm.io/gorm"
)

type QueryAdapter interface {
	CreateUser(db *gorm.DB, product *UserEntity) (UserEntity, error)
	CreateRole(db *gorm.DB, role *RoleEntity) (RoleEntity, error)
	CreatePermission(db *gorm.DB, permission *PermissionEntity) (PermissionEntity, error)
	GetUserByEmail(db *gorm.DB, email string) (UserEntity, error)
	GetRoleByID(db *gorm.DB, id uint) (RoleEntity, error)
}
