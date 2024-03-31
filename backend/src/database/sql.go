package database

import (
	"database/sql"
	"errors"

	"github.com/joshtyf/flowforge/src/logger"
)

var (

	// User
	CreateUserStatement = `INSERT INTO public."user" (user_id, name, identity_provider) 
								VALUES ($1, $2, $3) RETURNING created_on`

	SelectUserByIdStatement = `SELECT user_id, name, identity_provider, created_on, deleted 
								FROM public."user"
								WHERE user_id = $1
								AND deleted = false`

	CheckUserExistsStatement = `SELECT * 
								FROM public."user" 
								WHERE user_id = $1
								AND deleted = false`

	// Organisation
	CreateOrganizationStatement = `INSERT INTO public."organization" (name, owner) 
									VALUES ($1, $2) RETURNING org_id, created_on`

	SelectOrganizationsStatement = `SELECT o.* FROM public."organization" o
									INNER JOIN public."membership" m
									ON o.org_id = m.org_id
									WHERE user_id = $1
									AND o.deleted = false`

	SelectOrganisationByUserAndOrgIdStatement = `SELECT * FROM public."organization"
													WHERE org_id = $1
													AND owner = $2
													AND deleted = false`

	// Membership
	CreateMembershipStatement = `INSERT INTO public."membership" (user_id, org_id, role) 
								  VALUES ($1, $2, $3) RETURNING joined_on`

	SelectMembershipByUserAndOrgIdStatement = `SELECT * FROM public."membership"
												WHERE org_id = $1
												AND user_id = $2
												AND deleted = false`

	UpdateMembershipStatement = `UPDATE public."membership"
									SET role = $1
									WHERE user_id = $2
									AND org_id = $3
									AND deleted = false`

	DeleteMembershipStatement = `UPDATE public."membership"
									SET deleted = $1
									WHERE user_id = $2
									AND org_id = $3`
)

func txnRollback(tx *sql.Tx) {
	err := tx.Rollback()
	if !errors.Is(err, sql.ErrTxDone) {
		logger.Error("[Rollback] Error on rollback", map[string]interface{}{"err": err})
	}
}
