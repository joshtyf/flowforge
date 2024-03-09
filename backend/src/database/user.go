package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/logger"
)

type User struct {
	c *sql.DB
}

func NewUser(c *sql.DB) *User {
	return &User{c: c}
}

func txnRollback(tx *sql.Tx) {
	err := tx.Rollback()
	if !errors.Is(err, sql.ErrTxDone) {
		logger.Error("[Rollback] Error on rollback", map[string]interface{}{"err": err})
	}
}

func (u *User) CreateUser(user *models.UserModel) (*models.UserModel, error) {
	err := u.c.QueryRow(CheckUserExistsStatement, user.Id).Scan(&user.CreatedOn)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		if err = u.c.QueryRow(CreateUserStatement, user.Id, user.Name, user.ConnectionType).Scan(&user.CreatedOn); err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (u *User) CreateOrganisation(org *models.OrganisationModel) (*models.OrganisationModel, error) {
	tx, err := u.c.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	defer txnRollback(tx)

	if err := tx.QueryRow(CreateOrganisationStatement, org.Name, org.CreatedBy).Scan(&org.Id); err != nil {
		return nil, err
	}

	if _, err := tx.Exec(CreateMembershipStatement, org.CreatedBy, org.Id); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return org, nil
}

func (u *User) CreateMembership(membership *models.MembershipModel) (*models.MembershipModel, error) {
	if err := u.c.QueryRow(CreateMembershipStatement, membership.UserId, membership.OrgId).Scan(&membership.JoinedOn); err != nil {
		return nil, err
	}
	return membership, nil
}

func (u *User) GetUserById(user_id string) (*models.UserModel, error) {
	um := &models.UserModel{}
	if err := u.c.QueryRow(SelectUserByIdStatement, user_id).Scan(&um.Id, &um.Name, &um.ConnectionType, &um.CreatedOn); err != nil {
		return nil, err
	}
	return um, nil
}

func (u *User) GetUserOrganisations(user_id string) ([]*models.OrganisationModel, error) {
	rows, err := u.c.Query(SelectOrganisationsStatement, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	oms := []*models.OrganisationModel{}
	for rows.Next() {
		om := &models.OrganisationModel{}
		if err := rows.Scan(&om.Id, &om.Name, &om.CreatedBy, &om.CreatedOn); err != nil {
			return nil, err
		}
		oms = append(oms, om)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return oms, nil
}
