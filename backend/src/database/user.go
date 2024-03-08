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
	err := u.c.QueryRow(SelectUserStatement, user.Id).Scan()
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return user, nil
	}

	_, err = u.c.Exec(CreateUserStatement, user.Id, user.Name, user.ConnectionType)
	if err != nil {
		return nil, err
	}

	return user, nil
}
