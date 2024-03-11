//go:build unit
// +build unit

package createuser_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wstiehler/zpecreateuser-api/internal/domain/createuser"
	config "github.com/wstiehler/zpecreateuser-api/internal/infrastructure/database"
)

func TestService_CreatePermission(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	err = config.AutoMigrateTables(db)
	assert.NoError(t, err)

	repo := createuser.NewRepository(db, createuser.MemorySqlAdapter{})
	service := createuser.NewService(repo)

	newUser := &createuser.UserEntity{
		Name:     "William Villani Stiehler",
		Email:    "will@gmail.com",
		Password: "1212",
		RoleId:   1,
	}

	user, _ := service.CreateUser(db, newUser)

	assert.NotNil(t, user)
	assert.Equal(t, "will@gmail.com", user.Email)

	user2, err := service.CreateUser(db, newUser)

	expectedErr := errors.New("email already exists")
	assert.Equal(t, expectedErr.Error(), err.Error())
	assert.Nil(t, nil, user2)
}
