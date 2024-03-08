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
	Role      string    `json:"role"`
	Email     string    `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *UserEntity) TableName() string {
	return "users"
}
