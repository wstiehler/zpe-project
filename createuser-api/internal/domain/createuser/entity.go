package createuser

import (
	"time"
)

type ModelsEntity struct {
	UserEntity *UserEntity
}

type UserEntity struct {
	Id        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RoleId    uint      `json:"role_id"`
}

func (u *UserEntity) TableName() string {
	return "users"
}
