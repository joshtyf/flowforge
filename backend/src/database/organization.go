package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/joshtyf/flowforge/src/database/models"
)

type Organization struct {
	c *sql.DB
}

func NewOrganization(c *sql.DB) *Organization {
	return &Organization{c: c}
}

func (o *Organization) Create(org *models.OrganizationModel) (*models.OrganizationModel, error) {
	tx, err := o.c.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	defer txnRollback(tx)

	if err := tx.QueryRow(CreateOrganizationStatement, org.Name, org.Owner).Scan(&org.OrgId, &org.CreatedOn); err != nil {
		return nil, err
	}

	if _, err := tx.Exec(CreateMembershipStatement, org.Owner, org.OrgId, models.Owner); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return org, nil
}

func (o *Organization) GetAllOrgsByUserId(user_id string) ([]*models.OrganizationModel, error) {
	rows, err := o.c.Query(SelectOrganizationsStatement, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	oms := []*models.OrganizationModel{}
	for rows.Next() {
		om := &models.OrganizationModel{}
		if err := rows.Scan(&om.OrgId, &om.Name, &om.Owner, &om.CreatedOn, &om.Deleted); err != nil {
			return nil, err
		}
		oms = append(oms, om)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return oms, nil
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

func (o *Organization) GetAllUsersByOrgId(orgId int) ([]*GetAllUsersByOrdIdResponse, error) {
	rows, err := o.c.Query(SelectAllUsersByOrgIdStatement, orgId)
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

func (o *Organization) GetOrgByOrgIdAndOwner(user_id string, org_id int) (*models.OrganizationModel, error) {
	om := &models.OrganizationModel{}
	if err := o.c.QueryRow(SelectOrganizationByOrgIdAndOwnerStatement, org_id, user_id).Scan(&om.OrgId, &om.Name, &om.Owner, &om.CreatedOn, &om.Deleted); err != nil {
		return nil, err
	}
	return om, nil
}

func (m *Organization) UpdateOrgName(new_name string, org_id int, user_id string) (sql.Result, error) {
	result, err := m.c.Exec(UpdateOrganizationNameByOrgIdAndOwnerStatement, new_name, org_id, user_id)
	if err != nil {
		return nil, err
	}
	// NOTE: may not work for all db / db drivers
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, errors.New("unable to retrieve rows affected")
	} else if rows < 1 {
		return nil, fmt.Errorf("cannot find organization with org_id: %d and org_owner: %v", org_id, user_id)
	}
	return result, nil
}

func (o *Organization) DeleteOrg(org_id int, user_id string) (sql.Result, error) {

	result, err := o.c.Exec(DeleteOrganizationByOrgIdAndOwnerStatement, org_id, user_id)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, errors.New("unable to retrieve rows affected")
	} else if rows < 1 {
		return nil, fmt.Errorf("cannot find organization with org_id %d and org_owner %v", org_id, user_id)
	}
	return result, nil
}
