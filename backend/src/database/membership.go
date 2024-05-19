package database

import (
	"database/sql"
	"errors"

	"github.com/joshtyf/flowforge/src/database/models"
)

type Membership struct {
	c *sql.DB
}

func NewMembership(c *sql.DB) *Membership {
	return &Membership{c: c}
}

func (m *Membership) Create(membership *models.MembershipModel) (*models.MembershipModel, error) {
	if err := m.c.QueryRow(CreateMembershipStatement, membership.UserId, membership.OrgId, membership.Role).Scan(&membership.JoinedOn); err != nil {
		return nil, err
	}
	return membership, nil
}

func (m *Membership) GetUserMemberships(user_id string) ([]*models.MembershipModel, error) {
	rows, err := m.c.Query(GetUserMembershipsStatement, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	memberships := make([]*models.MembershipModel, 0)
	for rows.Next() {
		mm := &models.MembershipModel{}
		if err := rows.Scan(
			&mm.UserId,
			&mm.OrgId,
			&mm.Role,
			&mm.JoinedOn,
			&mm.Deleted,
		); err != nil {
			return nil, err
		}
		memberships = append(memberships, mm)
	}
	return memberships, nil
}

func (m *Membership) GetMembershipByUserAndOrgId(user_id string, org_id int) (*models.MembershipModel, error) {
	mm := &models.MembershipModel{}
	if err := m.c.QueryRow(SelectMembershipByUserAndOrgIdStatement, org_id, user_id).Scan(&mm.UserId, &mm.OrgId, &mm.Role, &mm.JoinedOn, &mm.Deleted); err != nil {
		return nil, err
	}
	return mm, nil
}

func (m *Membership) UpdateUserMembership(membership *models.MembershipModel) (sql.Result, error) {
	result, err := m.c.Exec(UpdateMembershipStatement, membership.Role, membership.UserId, membership.OrgId)
	if err != nil {
		return nil, err
	}
	// NOTE: may not work for all db / db drivers
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, errors.New("unable to retrieve rows affected")
	} else if rows < 1 {
		return nil, errors.New("membership does not exist")
	}
	return result, nil
}

func (m *Membership) DeleteUserMembership(membership *models.MembershipModel) (sql.Result, error) {
	result, err := m.c.Exec(DeleteMembershipStatement, true, membership.UserId, membership.OrgId)
	if err != nil {
		return nil, err
	}
	// NOTE: may not work for all db / db drivers
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, errors.New("unable to retrieve rows affected")
	} else if rows < 1 {
		return nil, errors.New("membership does not exist")
	}
	return result, nil
}
