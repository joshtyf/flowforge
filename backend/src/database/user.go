package database

import (
	"database/sql"
	"time"

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
	if err := u.c.QueryRow(SelectUserByIdStatement, user_id).Scan(&um.UserId, &um.Name, &um.IdentityProvider, &um.CreatedOn, &um.Deleted); err != nil {
		return nil, err
	}
	return um, nil
}
func (u *User) GetAllUsers() ([]*models.UserModel, error) {
	rows, err := u.c.Query(SelectAllUsersStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*models.UserModel, 0)
	for rows.Next() {
		um := &models.UserModel{}
		if err := rows.Scan(&um.UserId, &um.Name, &um.IdentityProvider, &um.CreatedOn, &um.Deleted); err != nil {
			return nil, err
		}
		users = append(users, um)
	}
	return users, nil
}

type GetAllUsersByOrdIdResponse struct {
	UserId           string    `json:"user_id"`
	Name             string    `json:"name"`
	IdentityProvider string    `json:"identity_provider"`
	CreatedOn        time.Time `json:"created_on"`
	Deleted          bool      `json:"deleted"`
	Role             string    `json:"role"`
	JoinedOn         time.Time `json:"joined_on"`
}

func (u *User) GetAllUsersByOrgId(orgId int) ([]*GetAllUsersByOrdIdResponse, error) {
	rows, err := u.c.Query(SelectAllUsersByOrgIdStatement, orgId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*GetAllUsersByOrdIdResponse, 0)
	for rows.Next() {
		um := &GetAllUsersByOrdIdResponse{}
		if err := rows.Scan(&um.UserId, &um.Name, &um.IdentityProvider, &um.CreatedOn, &um.Deleted, &um.Role, &um.JoinedOn); err != nil {
			return nil, err
		}
		users = append(users, um)
	}
	return users, nil
}
