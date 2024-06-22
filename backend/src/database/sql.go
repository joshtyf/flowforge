package database

import (
	"database/sql"
)

var (

	// User
	CreateUserStatement = `INSERT INTO public."user" (user_id, name, email, identity_provider) 
								VALUES ($1, $2, $3, $4) RETURNING created_on`

	SelectUserByIdStatement = `SELECT user_id, name, email, identity_provider, created_on, deleted 
								FROM public."user"
								WHERE user_id = $1
								AND deleted = false`

	SelectAllUsersStatement = `SELECT user_id, name, identity_provider, created_on, deleted 
								FROM public."user" WHERE deleted = false`

	SelectAllUsersByOrgIdStatement = `SELECT u.user_id, u.name, u.identity_provider, u.created_on, u.deleted, m.role, m.joined_on
										FROM public."user" u
										INNER JOIN public."membership" m
										ON u.user_id = m.user_id
										WHERE m.org_id = $1
										AND u.deleted = false
										AND m.deleted = false`

	CheckUserExistsStatement = `SELECT * 
								FROM public."user" 
								WHERE user_id = $1
								AND deleted = false`

	// Organization
	CreateOrganizationStatement = `INSERT INTO public."organization" (name, owner) 
									VALUES ($1, $2) RETURNING org_id, created_on`

	SelectOrganizationsStatement = `SELECT o.* FROM public."organization" o
									INNER JOIN public."membership" m
									ON o.org_id = m.org_id
									WHERE user_id = $1
									AND o.deleted = false`

	SelectOrganizationByOrgIdAndOwnerStatement = `SELECT * FROM public."organization"
													WHERE org_id = $1
													AND owner = $2
													AND deleted = false`

	UpdateOrganizationNameByOrgIdAndOwnerStatement = `UPDATE public."organization"
										SET name = $1
										WHERE org_id = $2
										AND owner = $3
										AND deleted = false`

	DeleteOrganizationByOrgIdAndOwnerStatement = `UPDATE public."organization"
															SET deleted = true
															WHERE org_id = $1
															AND owner = $2
															AND deleted = false`

	// Membership
	CreateMembershipStatement = `INSERT INTO public."membership" (user_id, org_id, role) 
								  VALUES ($1, $2, $3) RETURNING joined_on`

	RenewMembershipStatement = `UPDATE public."membership" 
								SET role = $3, joined_on = NOW(), deleted = false
								WHERE user_id = $1 AND org_id = $2
								RETURNING joined_on`

	GetUserMembershipsStatement = `SELECT * FROM public."membership"
									WHERE user_id = $1
									AND deleted = false`

	CheckMembershipRecordExistsStatement = `SELECT * FROM public."membership" WHERE user_id = $1 AND org_id = $2`

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

// TODO: figure out how to log this
func txnRollback(tx *sql.Tx) {
	tx.Rollback()
	// err := tx.Rollback()
	// if !errors.Is(err, sql.ErrTxDone) {
	// logger.Error("[Rollback] Error on rollback", map[string]interface{}{"err": err})
	// }
}
