package createuser

import (
	"time"
)

type ModelsEntity struct {
	UserEntity *UserEntity
	RoleEntity *RoleEntity
	Permission *PermissionEntity
}

type UserEntity struct {
	Id        string     `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"unique"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	RoleId    uint       `json:"role_id"`
	Role      RoleEntity `json:"role" gorm:"foreignKey:RoleId"`
}

type RoleEntity struct {
	ID          uint               `gorm:"primary_key"`
	Role        string             `json:"role" gorm:"unique"`
	Permissions []PermissionEntity `json:"permissions" gorm:"foreignKey:RoleId"`
}

type PermissionEntity struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	RoleId uint   `json:"role_id"`
	Name   string `json:"name" gorm:"unique"`
}

func (u *UserEntity) TableName() string {
	return "users"
}

func (r *RoleEntity) TableName() string {
	return "roles"
}

func (p *PermissionEntity) TableName() string {
	return "permissions"
}
