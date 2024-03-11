package role

import (
	"strings"

	"github.com/wstiehler/role-api/internal/infrastructure/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	repo Repository
}

func NewService(repo *Repository) *Service {
	return &Service{*repo}
}

func NormalizeString(s string) string {
	return strings.ToLower(s)
}

func (s *Service) CreateRole(db *gorm.DB, role *RoleEntity) (*RoleEntity, error) {
	role.Role = NormalizeString(role.Role)
	createdRole, err := s.repo.CreateRole(role)
	if err != nil {
		return nil, err
	}

	return &createdRole, nil
}

func (s *Service) CreatePermission(db *gorm.DB, permission *PermissionEntity) (*PermissionEntity, error) {
	permission.Name = NormalizeString(permission.Name)
	createdPermission, err := s.repo.CreatePermission(permission)
	if err != nil {
		return nil, err
	}

	return &createdPermission, nil
}

func (s *Service) GetRoleByID(id uint, db *gorm.DB) (*RoleDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	role, err := s.repo.GetRoleByID(id)
	if err != nil {
		logger.Error("Error on get user by email", zap.String("error", err.Error()))
		return nil, err
	}

	roleResponse := RoleDTO{
		Role: NormalizeString(role.Role),
		Id:   role.ID,
	}

	logger.Debug("Successfull on get role by id", zap.String("role", roleResponse.Role))

	return &roleResponse, nil
}

func (s *Service) GetPermissionsByRoleID(roleID uint, db *gorm.DB) ([]PermissionDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	permissions, err := s.repo.GetPermissionByID(roleID)
	if err != nil {
		logger.Error("Error retrieving permissions by role ID", zap.Error(err))
		return nil, err
	}

	var permissionDTOs []PermissionDTO
	for _, p := range permissions {
		permissionDTO := PermissionDTO{
			Name: p.Name,
		}
		permissionDTOs = append(permissionDTOs, permissionDTO)
	}

	logger.Debug("Successfully retrieved permissions by role ID", zap.Uint("roleID", roleID))

	return permissionDTOs, nil
}
