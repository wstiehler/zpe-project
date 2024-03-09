package role

import (
	"gorm.io/gorm"
)

type MysqlAdapter struct {
}

func (MysqlAdapter) CreateRole(db *gorm.DB, role *RoleEntity) (RoleEntity, error) {
	if err := db.Create(role).Error; err != nil {
		return RoleEntity{}, err
	}
	return *role, nil
}

func (MysqlAdapter) CreatePermission(db *gorm.DB, permission *PermissionEntity) (PermissionEntity, error) {
	if err := db.Create(permission).Error; err != nil {
		return PermissionEntity{}, err
	}
	return *permission, nil
}

func (MysqlAdapter) GetRoleByID(db *gorm.DB, id uint) (RoleEntity, error) {
	var role RoleEntity
	db.Where("id = ?", id).First(&role)

	return role, nil
}

func (MysqlAdapter) GetPermissionByID(db *gorm.DB, id uint) ([]PermissionEntity, error) {
	var permissions []PermissionEntity
	if err := db.Where("role_id = ?", id).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}
