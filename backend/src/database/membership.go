package database

import (
	"database/sql"

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
	return result, nil
}

func (m *Membership) DeleteUserMembership(membership *models.MembershipModel) (sql.Result, error) {
	result, err := m.c.Exec(DeleteMembershipStatement, true, membership.UserId, membership.OrgId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
