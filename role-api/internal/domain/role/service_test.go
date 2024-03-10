package role_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wstiehler/role-api/internal/domain/role"
	config "github.com/wstiehler/role-api/internal/infrastructure/database"
)

func TestService_CreateRole(t *testing.T) {
	// Conectar ao banco de dados em memória
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)
	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	// Criar as tabelas
	err = config.AutoMigrateTables(db)
	assert.NoError(t, err)

	// Criar o repositório
	repo := role.NewRepository(db, role.MemorySqlAdapter{})

	// Criar o serviço
	service := role.NewService(repo)

	// Definir o novo papel a ser criado
	newRole := &role.RoleEntity{
		ID:   3,
		Role: "Editor",
	}

	// Criar o papel
	createdRole, _ := service.CreateRole(db, newRole)

	// Verificar se o papel criado não é nulo
	assert.NotNil(t, createdRole)

	// Verificar se o nome do papel criado foi normalizado para minúsculas
	assert.Equal(t, "editor", createdRole.Role)

	// Tentar criar um papel com o mesmo ID novamente
	_, err = service.CreateRole(db, newRole)

	// Verificar se ocorreu um erro esperado
	assert.Error(t, err)
	// Verificar se a mensagem de erro é a esperada
	assert.Equal(t, "UNIQUE constraint failed: roles.id", err.Error())
}
