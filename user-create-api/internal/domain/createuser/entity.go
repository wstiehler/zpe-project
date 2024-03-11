package createuser

import (
	"time"
)

type ModelsEntity struct {
	UserEntity *UserEntity
}

type UserEntity struct {
	Id        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RoleId    uint      `json:"role_id"`
}

type RoleEntity struct {
	ID          uint               `gorm:"primary_key"`
	Role        string             `json:"role" gorm:"unique"`
	Permissions []PermissionEntity `json:"permissions" gorm:"foreignKey:RoleId"`
}

type PermissionEntity struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	RoleId uint   `json:"role_id" gorm:"index:idx_role_id"`
	Name   string `json:"name"`
}

func (u *RoleEntity) TableName() string {
	return "roles"
}

func (u *UserEntity) TableName() string {
	return "users"
}
