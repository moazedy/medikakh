package queries

const (
	InsertDataQuery           = `INSERT INTO dd (KEY, VALUE) VALUES ($key, $value)`
	ReadDataByIdQuery         = `SELECT dd.* FROM dd WHERE meta().id= $1`
	ReadDataByTitleQuery      = ` SELECT dd.* FROM dd WHERE title=$1`
	ReadDataUsingPatternQuery = `SELECT dd.title FROM dd WHERE title  LIKE $1`
	IsDDExistsQuery           = "SELECT meta().id from dd WHERE dd.title = $1"
)
