package queries

const (
	RegisterUserQuery = `INSERT INTO users (KEY,VALUE) ` +
		` VALUES ($id, $user)`

	ReadUserByIdQuery = `SELECT * FROM users WHERE meta().id = $1`

	ReadUserByUsernameQuery = `SELECT * FROM users WHERE username= $1`

	DeleteUserQuery = `DELETE FROM users WHERE meta().id= $1`

	GetUserPasswordQuery = `SELECT  users.password FROM users WHERE meta().id= $1`

	GetUserRoleQuery = `SELECT users.role FROM users WHERE meta().id = $1`

	GetUserIdByUsernameQuery = `SELECT meta().id FORM users WHERE users.username= $1`

	UpdateUserQuery = `UPDATE users SET username=$1 password=$2 role=$3 created_at=$4`

	IsUsernameExistsQuery = ` SELECT meta(),id FROM users WHERE users.username= $1`
)
