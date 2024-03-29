//go:build unit
// +build unit

package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wstiehler/zpeupdateuser-api/internal/domain/user"
	config "github.com/wstiehler/zpeupdateuser-api/internal/infrastructure/database"
)

func TestService_UpdateUser(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	repo := user.NewRepository(db, user.MemorySqlAdapter{})
	service := user.NewService(repo)

	u, _ := repo.GetUserByEmail("will@gmail.com")

	UpdateUser := &user.UserEntity{
		Name: "William Villani",
	}

	user, _ := service.UpdateUser(u.Id, *UpdateUser, db)

	assert.NotNil(t, user)
	assert.Equal(t, "William Villani", user.Name)

}

func TestService_DeleteUser(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	repo := user.NewRepository(db, user.MemorySqlAdapter{})
	service := user.NewService(repo)

	u, _ := repo.GetUserByEmail("will@gmail.com")

	err = service.DeleteUser(u.Id, db)

	assert.Nil(t, err)

}

func TestService_GetUserInformation_InvalidCredentials(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	repo := user.NewRepository(db, user.MemorySqlAdapter{})
	service := user.NewService(repo)

	email := "will@gmail.com"
	password := "invalidpassword"

	_, _, err = service.GetUserInformation(email, password)
	assert.Error(t, err)
}

func TestService_CreateJWToken(t *testing.T) {
	service := user.NewService(&user.Repository{})

	userID := "123"
	role := "admin"
	permissions := []string{"read", "write"}

	token, err := service.CreateJWToken(userID, role, permissions)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestService_Login_InvalidCredentials(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	repo := user.NewRepository(db, user.MemorySqlAdapter{})
	service := user.NewService(repo)

	userLogin := user.UserLoginEntity{
		Email:    "invalid@example.com",
		Password: "invalidpassword",
	}
	_, err = service.Login(db, userLogin)

	assert.Error(t, err)
}

func TestNormalizeString(t *testing.T) {
	normalized := user.NormalizeString("Hello")
	assert.Equal(t, "hello", normalized)

	normalized = user.NormalizeString("HeLLo")
	assert.Equal(t, "hello", normalized)
}
