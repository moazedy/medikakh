package queries

const (
	SaveArticleQuery = `INSERT INTO articles (KEY,VALUE) VALUES ($id, $article)`

	ReadArticleByIdQuery = `SELECT * FROM articles WHERE meta().id = $id`

	ReadArticleByTitleQuery = `SELECT * FORM articles WHERE articles.title = $title`

	GetArticleCategoryQuery = `SELECT articles.category FROM articles WHERE meta().id = $1`

	GetArticleSubCategoryQuery = `SELECT articles.sub_category FROM articles WHRE meta().id=$1`

	GetArticleIdQuery = `SELECT articles.Id FORM articles WHERE aritcles.title = $1`

	ReadArticleSummeryQuery = `SELECT articles.summery FROM articles WHRE meta().id=$1`

	ReadArticleContentQuery = `SELECT articles.content FROM articles WHRE meta().id=$1`

	ReadArticleResultQuery = `SELECT articles.result FROM articles WHRE meta().id=$1`

	DeleteArticleQuery = `DELETE FROM articles WHERE meta().id= $1`

	GetArticleStatusQuery = `SELECT articles.status FROM articles WHERE meta().id = $1`

	IsArticleExistsQuery = `SELECT meta().id FROM articles WHERE articles.title = $1`

	UpdateArticleQuery = `UPDATE articles SET title=$1 status=$2 summery=$3 content=$4 result=$5 category=$6 sub_category=$7 WHERE meta().id = $8`
)
