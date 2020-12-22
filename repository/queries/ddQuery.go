package queries

const (
	InsertDataQuery           = `INSERT INTO dd (KEY, VALUE) VALUES ($key, $value)`
	ReadDataByIdQuery         = `SELECT * FROM dd WHERE meta().id= $1`
	ReadDataByTitleQuery      = ` SELECT * FROM dd WHERE dd.title=$1`
	ReadDataUsingPatternQuery = `select * from dd where title  like $1`
)
