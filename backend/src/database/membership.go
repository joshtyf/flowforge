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

func (u *Membership) Create(membership *models.MembershipModel) (*models.MembershipModel, error) {
	if err := u.c.QueryRow(CreateMembershipStatement, membership.UserId, membership.OrgId).Scan(&membership.JoinedOn); err != nil {
		return nil, err
	}
	return membership, nil
}
