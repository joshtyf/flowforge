package database

var (
	CreateUserStatement = `INSERT INTO user (user_id, name, connection_type) VALUES ($1, $2, $3)`
	SelectUserStatement = `SELECT * FROM user WHERE user_id = $1`
)
