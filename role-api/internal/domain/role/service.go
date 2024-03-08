package role

import (
	"strings"

	"gorm.io/gorm"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
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
