package createuser

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

func (r *Repository) CreateUser(user *UserEntity) (UserEntity, error) {
	return r.queryAdapter.CreateUser(r.db, user)
}

func (u *Repository) GetUserByEmail(email string) (UserEntity, error) {
	return u.queryAdapter.GetUserByEmail(u.db, email)
}

func (u *Repository) GetRoleByName(role string) (RoleEntity, error) {
	return u.queryAdapter.GetRoleByName(u.db, role)
}
