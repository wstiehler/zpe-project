package user

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
	return u.queryAdapter.GetUserById(u.db, id)
}

func (u *Repository) UpdateUser(user *UserEntity, id string) (err error) {
	return u.queryAdapter.UpdateUser(u.db, user, id)
}

func (c *Repository) DeleteUser(id string) error {
	return c.queryAdapter.DeleteUser(c.db, id)
}
