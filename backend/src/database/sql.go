package database

var (

	// User
	CreateUserStatement      = `INSERT INTO public."user" (user_id, name, connection_type) VALUES ($1, $2, $3) RETURNING created_at`
	SelectUserByIdStatement  = `SELECT user_id, name, connection_type, created_at FROM public."user" WHERE user_id = $1`
	CheckUserExistsStatement = `SELECT created_at FROM public."user" WHERE user_id = $1`

	// Organisation
	CreateOrganisationStatement  = `INSERT INTO public."organisation" (name, created_by) VALUES ($1, $2) RETURNING org_id`
	SelectOrganisationsStatement = `SELECT org_id, name, created_by, created_at FROM public."organisation" WHERE org_id IN (SELECT org_id FROM public."membership" WHERE user_id = $1)`

	// Membership
	CreateMembershipStatement = `INSERT INTO public."membership" (user_id, org_id) VALUES ($1, $2) RETURNING joined_at`
)