package queries

const (
	SaveArticleQuery = `INSERT INTO articles (KEY,VALUE) VALUES ($id, $article)`

	ReadArticleByIdQuery = `SELECT articles.* FROM articles WHERE meta().id = $id`

	ReadArticleByTitleQuery = `SELECT articles.* FROM articles WHERE title = $title`

	GetArticleCategoryQuery = `SELECT articles.category FROM articles WHERE meta().id = $1`

	GetArticleSubCategoryQuery = `SELECT articles.sub_category FROM articles WHRE meta().id=$1`

	GetArticleIdQuery = `SELECT articles.Id FROM articles WHERE title = $1`

	ReadArticleSummeryQuery = `SELECT articles.summery FROM articles WHRE meta().id=$1`

	ReadArticleClinicalFeaturesQuery = `SELECT articles.clinical_features FROM articles WHRE meta().id=$1`

	ReadArticleEtiologyQuery = `SELECT articles.etiology FROM articles WHRE meta().id=$1`

	DeleteArticleQuery = `DELETE FROM articles WHERE meta().id= $1`

	GetArticleStatusQuery = `SELECT articles.status FROM articles WHERE meta().id = $1`

	IsArticleExistsQuery = `SELECT meta().id FROM articles WHERE articles.title = $1`

	UpdateArticleQuery = `UPDATE articles SET title=$1, status=$2, summary=$3, etiology=$4, clinical_featuers=$5, ` +
		` diagnostics=$6, treatment=$7, complications=$8, prevention=$9, references=$10, category=$11, sub_category=$12 ` +
		` WHERE meta().id = $13`

	ReadArticleDiagnostics = `SELECT articles.diagnostics FROM articles WHERE meta().id = $1`

	ReadArticleTreatment = `SELECT articles.treatment FROM articles WHERE meta().id = $1`

	ReadArticleComplicationsQuery = `SELECT articles.complications FROM articles WHERE meta().id = $1`

	ReadArticlePreventionQuery = `SELECT articles.prevention FROM articles WHERE meta().id = $1`

	ReadArticleReferencesQeury = `SELECT articles.references FROM articles WHERE meta().id = $1`

	GetArticlesByCategoreyQuery = `SELECT	articles.title FROM articles WHERE articles.category=$1`

	GetAllArticlesQuery = `SELECT	articles.title FROM articles `

	GetArticlesBySubCategoryQuery = `SELECT	articles.title FROM articles WHERE articles.category=$1 AND articles.sub_category=$2`

	GetArticlesTitleListQuery = "SELECT articles.title FROM articles"
)
