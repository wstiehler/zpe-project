package role_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wstiehler/role-api/internal/domain/role"
	config "github.com/wstiehler/role-api/internal/infrastructure/database"
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

	repo := role.NewRepository(db, role.MemorySqlAdapter{})
	service := role.NewService(repo)

	newPermission := &role.PermissionEntity{
		RoleId: 1,
		Name:   "Create",
	}

	createdPermission, _ := service.CreatePermission(db, newPermission)

	assert.NotNil(t, createdPermission)
	assert.Equal(t, "create", createdPermission.Name)
}

func TestService_GetRoleByID(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	err = config.AutoMigrateTables(db)
	assert.NoError(t, err)

	repo := role.NewRepository(db, role.MemorySqlAdapter{})
	service := role.NewService(repo)

	newRole := &role.RoleEntity{
		Role: "xpto",
	}

	_, err = service.CreateRole(db, newRole)
	assert.NoError(t, err)

	roleResponse, err := service.GetRoleByID(1, db)

	assert.NoError(t, err)
	assert.NotNil(t, roleResponse)
	assert.Equal(t, "xpto", roleResponse.Role)
}

func TestService_GetPermissionsByRoleID(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	err = config.AutoMigrateTables(db)
	assert.NoError(t, err)

	repo := role.NewRepository(db, role.MemorySqlAdapter{})
	service := role.NewService(repo)

	newPermission := &role.PermissionEntity{
		RoleId: 1,
		Name:   "Create",
	}

	_, err = service.CreatePermission(db, newPermission)
	assert.NoError(t, err)

	permissions, err := service.GetPermissionsByRoleID(1, db)

	assert.NoError(t, err)
	assert.NotNil(t, permissions)
	assert.Equal(t, "create", permissions[0].Name)
}
