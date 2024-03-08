package createuser

import (
	"gorm.io/gorm"
)

type MysqlAdapter struct {
}

func (MysqlAdapter) CreateUser(db *gorm.DB, user *UserEntity) (UserEntity, error) {
	if err := db.Create(user).Error; err != nil {
		return UserEntity{}, err
	}
	return *user, nil
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

func (MysqlAdapter) GetUserByEmail(db *gorm.DB, email string) (UserEntity, error) {
	var user UserEntity
	db.Where("email = ?", email).First(&user)
	return user, nil
}

func (MysqlAdapter) GetRoleByID(db *gorm.DB, id uint) (RoleEntity, error) {
	var role RoleEntity
	db.Where("id = ?", id).First(&role)

	return role, nil
}
