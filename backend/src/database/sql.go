package database

var (

	// User
	CreateUserStatement = `INSERT INTO public."user" (user_id, name, connection_type) VALUES ($1, $2, $3) RETURNING created_at`
	SelectUserStatement = `SELECT * FROM public."user" WHERE user_id = $1`

	// Organisation
	CreateOrganisationStatement = `INSERT INTO public."organisation" (name, created_by) VALUES ($1, $2) RETURNING org_id`

	// Membership
	CreateMembershipStatement = `INSERT INTO public."membership" (user_id, org_id) VALUES ($1, $2) RETURNING joined_at`
)
