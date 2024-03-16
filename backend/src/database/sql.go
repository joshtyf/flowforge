package database

import (
	"database/sql"
	"errors"

	"github.com/joshtyf/flowforge/src/logger"
)

var (

	// User
	CreateUserStatement = `INSERT INTO public."user" (user_id, name, connection_type) 
								VALUES ($1, $2, $3) RETURNING created_at`

	SelectUserByIdStatement = `SELECT user_id, name, connection_type, created_at 
								FROM public."user" 
								WHERE user_id = $1`

	CheckUserExistsStatement = `SELECT created_at 
								FROM public."user" 
								WHERE user_id = $1`

	// Organisation
	CreateOrganizationStatement = `INSERT INTO public."organization" (name, owner) 
									VALUES ($1, $2) RETURNING org_id`

	SelectOrganizationsStatement = `SELECT o.* FROM public."organization" o
									INNER JOIN public."membership" m
									ON o.org_id = m.org_id
									WHERE user_id = $1`

	// Membership
	CreateMembershipStatement = `INSERT INTO public."membership" (user_id, org_id) 
								  VALUES ($1, $2) RETURNING user_id, org_id, joined_at`
)

func txnRollback(tx *sql.Tx) {
	err := tx.Rollback()
	if !errors.Is(err, sql.ErrTxDone) {
		logger.Error("[Rollback] Error on rollback", map[string]interface{}{"err": err})
	}
}
