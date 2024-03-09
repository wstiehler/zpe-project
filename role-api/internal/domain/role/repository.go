package role

import "gorm.io/gorm"

type Repository struct {
	db           *gorm.DB
	queryAdapter QueryAdapter
}

func NewRepository(db *gorm.DB, queryAdapter QueryAdapter) *Repository {
	return &Repository{
		db:           db,
		queryAdapter: queryAdapter,
	}
}

func (r *Repository) CreateRole(role *RoleEntity) (RoleEntity, error) {
	return r.queryAdapter.CreateRole(r.db, role)
}

func (r *Repository) CreatePermission(permission *PermissionEntity) (PermissionEntity, error) {
	return r.queryAdapter.CreatePermission(r.db, permission)
}

func (u *Repository) GetRoleByID(id uint) (RoleEntity, error) {
	return u.queryAdapter.GetRoleByID(u.db, id)
}

func (u *Repository) GetPermissionByID(id uint) ([]PermissionEntity, error) {
	return u.queryAdapter.GetPermissionByID(u.db, id)
}
