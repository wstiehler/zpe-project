package detailuser

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

func (u *Repository) GetUserByEmail(email string) (UserEntity, error) {
	return u.queryAdapter.GetUserByEmail(u.db, email)
}

func (u *Repository) GetUserByID(id string) (UserEntity, error) {
	return u.queryAdapter.GetUserByID(u.db, id)
}

func (u *Repository) GetUserByName(name string) (UserEntity, error) {
	return u.queryAdapter.GetUserByName(u.db, name)
}

func (c *Repository) GetAllUsers() ([]UserEntity, error) {
	return c.queryAdapter.GetAllUsers(c.db)
}
