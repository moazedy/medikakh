package queries

const (
	InsertCategoryQuery = `INSERT INTO categories (KEY, VALUE) VALUES ($key, $category)`

	ReadCategoryByIdQuery = `SELECT * FROM categories WHERE meta().id= $1`

	ReadCategoryByNameQuery = `SELECT * FROM categories WHERE categories.name= $1`

	ReadCategorySubCategoriesQuery = `SELECT categories.sub_category FROM categories ` +
		` WHERE meta().id=$1`

	GetCategoryIdQuery = `SELECT categories.id FROM categories WHERE categories.name = $1`

	GetCategoriesQuery = `SELECT * FROM categories`
)
