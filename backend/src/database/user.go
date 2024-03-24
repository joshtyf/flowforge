package database

import (
	"database/sql"

	"github.com/joshtyf/flowforge/src/database/models"
)

type User struct {
	c *sql.DB
}

func NewUser(c *sql.DB) *User {
	return &User{c: c}
}

func (u *User) Create(user *models.UserModel) (*models.UserModel, error) {
	if err := u.c.QueryRow(CreateUserStatement, user.UserId, user.Name, user.IdentityProvider).Scan(&user.CreatedOn); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) GetUserById(user_id string) (*models.UserModel, error) {
	um := &models.UserModel{}
	if err := u.c.QueryRow(SelectUserByIdStatement, user_id).Scan(&um.UserId, &um.Name, &um.IdentityProvider, &um.CreatedOn); err != nil {
		return nil, err
	}
	return um, nil
}
