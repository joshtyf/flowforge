package database

import (
	"context"
	"database/sql"

	"github.com/joshtyf/flowforge/src/database/models"
)

type Organization struct {
	c *sql.DB
}

func NewOrganization(c *sql.DB) *Organization {
	return &Organization{c: c}
}

func (o *Organization) Create(org *models.OrganisationModel) (*models.OrganisationModel, error) {
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

func (o *Organization) GetAllOrgsByUserId(user_id string) ([]*models.OrganisationModel, error) {
	rows, err := o.c.Query(SelectOrganizationsStatement, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	oms := []*models.OrganisationModel{}
	for rows.Next() {
		om := &models.OrganisationModel{}
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

func (o *Organization) GetOrgByOwnerAndOrgId(user_id string, org_id int) (*models.OrganisationModel, error) {
	om := &models.OrganisationModel{}
	if err := o.c.QueryRow(SelectOrganisationByUserAndOrgId, org_id, user_id).Scan(&om.OrgId, &om.Name, &om.Owner, &om.CreatedOn, &om.Deleted); err != nil {
		return nil, err
	}
	return om, nil
}
