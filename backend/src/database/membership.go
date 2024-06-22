package database

import (
	"context"
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
	row := m.c.QueryRow(CheckMembershipRecordExistsStatement, membership.UserId, membership.OrgId)
	sqlStatementToExecute := CreateMembershipStatement

	// If membership record previously existed, renew membership
	if err := row.Scan(); err != sql.ErrNoRows {
		sqlStatementToExecute = RenewMembershipStatement
	}

	if err := m.c.QueryRow(sqlStatementToExecute, membership.UserId, membership.OrgId, membership.Role).Scan(&membership.JoinedOn); err != nil {
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

func (m *Membership) TransferOwnership(owner *models.MembershipModel, newOwner *models.MembershipModel) error {
	tx, err := m.c.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	defer txnRollback(tx)

	if _, err := tx.Exec(UpdateMembershipStatement, models.Admin, owner.UserId, owner.OrgId); err != nil {
		return err
	}

	if _, err := tx.Exec(UpdateMembershipStatement, models.Owner, newOwner.OrgId, newOwner.OrgId); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
