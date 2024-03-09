package user

import (
	"time"
)

type ModelsEntity struct {
	UserEntity *UserEntity
}

type UserEntity struct {
	Id        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" `
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RoleId    uint      `json:"role_id"`
}

type UserLoginEntity struct {
	Id       string `json:"id"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (u *UserEntity) TableName() string {
	return "users"
}
