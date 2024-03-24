package database

import (
	"database/sql"
)

var (

	// User
	CreateUserStatement = `INSERT INTO public."user" (user_id, name, identity_provider) 
								VALUES ($1, $2, $3) RETURNING created_on`

	SelectUserByIdStatement = `SELECT user_id, name, identity_provider, created_on 
								FROM public."user" 
								WHERE user_id = $1`

	CheckUserExistsStatement = `SELECT * 
								FROM public."user" 
								WHERE user_id = $1`

	// Organisation
	CreateOrganizationStatement = `INSERT INTO public."organization" (name, owner) 
									VALUES ($1, $2) RETURNING org_id, created_on`

	SelectOrganizationsStatement = `SELECT o.* FROM public."organization" o
									INNER JOIN public."membership" m
									ON o.org_id = m.org_id
									WHERE user_id = $1`

	// Membership
	CreateMembershipStatement = `INSERT INTO public."membership" (user_id, org_id) 
								  VALUES ($1, $2) RETURNING joined_on`
)

// TODO: figure out how to log this
func txnRollback(tx *sql.Tx) {
	tx.Rollback()
	// err := tx.Rollback()
	// if !errors.Is(err, sql.ErrTxDone) {
	// logger.Error("[Rollback] Error on rollback", map[string]interface{}{"err": err})
	// }
}
