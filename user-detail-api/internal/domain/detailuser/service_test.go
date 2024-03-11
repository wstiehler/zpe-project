//go:build unit
// +build unit

package detailuser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wstiehler/zpedetailuser-api/internal/domain/detailuser"
	config "github.com/wstiehler/zpedetailuser-api/internal/infrastructure/database"
)

func TestService_GetUserByEmail(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	repo := detailuser.NewRepository(db, detailuser.MemorySqlAdapter{})
	service := detailuser.NewService(repo)

	email := "will@gmail.com"
	_, err = service.GetUserByEmail(db, email)
	assert.NoError(t, err)

	_, err = service.GetUserByEmail(db, "nonexistent@example.com")
	assert.Error(t, err)
	assert.Equal(t, detailuser.ErrUserNotFound, err)
}

func TestService_GetUserByID(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	repo := detailuser.NewRepository(db, detailuser.MemorySqlAdapter{})
	service := detailuser.NewService(repo)

	id := "1"
	_, err = service.GetUserByID(db, id)
	assert.Equal(t, detailuser.ErrUserNotFound, err)

	_, err = service.GetUserByID(db, "nonexistent_id")
	assert.Error(t, err)
	assert.Equal(t, detailuser.ErrUserNotFound, err)
}

func TestService_GetUserByName(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	repo := detailuser.NewRepository(db, detailuser.MemorySqlAdapter{})
	service := detailuser.NewService(repo)

	name := "William Villani Stiehler"
	_, err = service.GetUserByName(db, name)
	assert.NoError(t, err)

	_, err = service.GetUserByName(db, "nonexistent_name")
	assert.Error(t, err)
	assert.Equal(t, detailuser.ErrUserNotFound, err)
}

func TestService_GetUserByCriteria(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	repo := detailuser.NewRepository(db, detailuser.MemorySqlAdapter{})
	service := detailuser.NewService(repo)

	email := "will@gmail.com"
	users, err := service.GetUserByCriteria(db, "email", email)
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.NotEmpty(t, users)

	nonexistentEmail := "nonexistent@example.com"
	_, err = service.GetUserByCriteria(db, "email", nonexistentEmail)
	assert.Nil(t, err)
}
