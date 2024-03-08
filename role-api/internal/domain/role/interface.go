package role

import (
	"gorm.io/gorm"
)

type QueryAdapter interface {
	CreateRole(db *gorm.DB, role *RoleEntity) (RoleEntity, error)
	CreatePermission(db *gorm.DB, permission *PermissionEntity) (PermissionEntity, error)
	GetRoleByID(db *gorm.DB, id uint) (RoleEntity, error)
}
