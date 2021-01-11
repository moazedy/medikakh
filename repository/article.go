package repository

import (
	"errors"
	"github.com/couchbase/gocb/v2"
	"medikakh/domain/models"
	"medikakh/repository/queries"
)

type ArticleRepo interface {
	Save(article models.Article) error
	ReadArticleById(id string) (*models.Article, error)
	ReadArticleByTitle(title string) (*models.Article, error)
	GetArticleCategory(id string) (*string, error)
	GetArticleSubsCategory(id string) (*string, error)
	GetArticleId(title string) (*string, error)
	ReadArticleSummary(id string) (*string, error)
	ReadArticleEtiology(id string) (*string, error)
	ReadArticleClinicalFeatures(id string) (*string, error)
	ReadArticleDiagnostics(id string) (*string, error)
	ReadArticleTreatment(id string) (*string, error)
	ReadArticleComplications(id string) (*string, error)
	ReadArticlePrevention(id string) (*string, error)
	ReadArticleReferences(id string) (*string, error)
	UpdateArticle(article models.Article) error
	DeleteArticle(id string) error
	GetArticleStatus(id string) (*string, error)
	IsArticleExists(title string) (*bool, error)
	GetArticlesByCategory(cat string) ([]models.Article, error)
	GetAllArticles() ([]models.Article, error)
	GetArticlesBySubCategory(cat, subCat string) ([]models.Article, error)
	GetArticlesList() ([]string, error)
}

type article struct {
	session *gocb.Cluster
}

func NewArticleRpo(session *gocb.Cluster) ArticleRepo {
	a := new(article)
	a.session = session
	return a
}

func (a *article) Save(article models.Article) error {
	_, err := a.session.Query(
		queries.SaveArticleQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"id":      article.Id,
			"article": article,
		}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *article) ReadArticleById(id string) (*models.Article, error) {
	res, err := a.session.Query(
		queries.ReadArticleByIdQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"id": id,
		}},
	)
	if err != nil {
		return nil, err
	}

	var newArticle models.Article
	for res.Next() {
		err = res.Row(&newArticle)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("the article does not exist !")
			}

			return nil, err
		}
	}

	return &newArticle, nil
}

func (a *article) ReadArticleByTitle(title string) (*models.Article, error) {
	res, err := a.session.Query(
		queries.ReadArticleByTitleQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"title": title,
		}},
	)
	if err != nil {
		return nil, err
	}

	var newArticle models.Article
	for res.Next() {
		err = res.Row(&newArticle)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("the article does not exist !")
			}

			return nil, err
		}
	}

	return &newArticle, nil

}

func (a *article) GetArticleCategory(id string) (*string, error) {
	res, err := a.session.Query(
		queries.GetArticleCategoryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var category string
	for res.Next() {
		err = res.Row(&category)
		if err != nil {
			return nil, err
		}
	}

	return &category, nil
}

func (a *article) GetArticleSubsCategory(id string) (*string, error) {
	res, err := a.session.Query(
		queries.GetArticleSubCategoryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var subCategory string
	for res.Next() {
		err = res.Row(&subCategory)
		if err != nil {
			return nil, err
		}
	}

	return &subCategory, nil
}

func (a *article) GetArticleId(title string) (*string, error) {
	res, err := a.session.Query(
		queries.GetArticleIdQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{title}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var id string
	for res.Next() {
		err = res.Row(&id)
		if err != nil {
			return nil, err
		}
	}

	return &id, nil

}
func (a *article) ReadArticleSummary(id string) (*string, error) {
	res, err := a.session.Query(
		queries.ReadArticleSummeryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var summery string
	for res.Next() {
		err = res.Row(&summery)
		if err != nil {
			return nil, err
		}
	}

	return &summery, nil

}

func (a *article) ReadArticleEtiology(id string) (*string, error) {
	res, err := a.session.Query(
		queries.ReadArticleEtiologyQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var result string
	for res.Next() {
		err = res.Row(&result)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil

}

func (a *article) ReadArticleClinicalFeatures(id string) (*string, error) {
	res, err := a.session.Query(
		queries.ReadArticleClinicalFeaturesQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var content string
	for res.Next() {
		err = res.Row(&content)
		if err != nil {
			return nil, err
		}
	}

	return &content, nil

}

func (a *article) DeleteArticle(id string) error {
	_, err := a.session.Query(
		queries.DeleteArticleQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *article) GetArticleStatus(id string) (*string, error) {
	res, err := a.session.Query(
		queries.GetArticleStatusQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var status string
	for res.Next() {
		err = res.Row(&status)
		if err != nil {
			return nil, err
		}
	}

	return &status, nil

}

func (a *article) IsArticleExists(title string) (*bool, error) {
	res, err := a.session.Query(
		queries.IsArticleExistsQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{title}},
	)
	if err != nil {
		return nil, errors.New("error on serching for specific article")
	}
	var returnValue bool
	var id string
	for res.Next() {
		err = res.Row(&id)
		if err != nil {
			if err == gocb.ErrNoResult {
				return &returnValue, errors.New("article does not exist") //it should be checked later
			}
			return nil, err
		}
	}

	if id != "" {
		returnValue = true
	}

	return &returnValue, nil
}

func (a *article) UpdateArticle(article models.Article) error {
	_, err := a.session.Query(
		queries.UpdateArticleQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{
			article.Title,
			article.Status,
			article.Summary,
			article.Etiology,
			article.ClinicalFeatures,
			article.Diagnostics,
			article.Treatment,
			article.Complications,
			article.Prevention,
			article.References,
			article.Category,
			article.SubCategory,
			article.Id,
		}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *article) ReadArticleDiagnostics(id string) (*string, error) {
	res, err := a.session.Query(
		queries.ReadArticleDiagnostics,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var diags string
	for res.Next() {
		err = res.Row(&diags)
		if err != nil {
			return nil, err
		}
	}

	return &diags, nil

}

func (a *article) ReadArticleTreatment(id string) (*string, error) {
	res, err := a.session.Query(
		queries.ReadArticleTreatment,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var treatments string
	for res.Next() {
		err = res.Row(&treatments)
		if err != nil {
			return nil, err
		}
	}

	return &treatments, nil

}

func (a *article) ReadArticleComplications(id string) (*string, error) {
	res, err := a.session.Query(
		queries.ReadArticleComplicationsQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var comps string
	for res.Next() {
		err = res.Row(&comps)
		if err != nil {
			return nil, err
		}
	}

	return &comps, nil

}

func (a *article) ReadArticlePrevention(id string) (*string, error) {
	res, err := a.session.Query(
		queries.ReadArticlePreventionQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var prevs string
	for res.Next() {
		err = res.Row(&prevs)
		if err != nil {
			return nil, err
		}
	}

	return &prevs, nil

}

func (a *article) ReadArticleReferences(id string) (*string, error) {
	res, err := a.session.Query(
		queries.ReadArticleReferencesQeury,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("the article does not exist !")
		}

		return nil, err
	}

	var refs string
	for res.Next() {
		err = res.Row(&refs)
		if err != nil {
			return nil, err
		}
	}

	return &refs, nil

}

func (a *article) GetArticlesByCategory(cat string) ([]models.Article, error) {
	res, err := a.session.Query(
		queries.GetArticlesByCategoreyQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{cat}},
	)
	if err != nil {
		return nil, errors.New("error on reading items from db")
	}

	var articles []models.Article
	for res.Next() {
		var art models.Article
		err = res.Row(&art)
		if err != nil {
			if err == gocb.ErrNoResult {
				return articles, nil
			}
			return nil, err
		}
		articles = append(articles, art)
	}

	return articles, nil
}

func (a *article) GetAllArticles() ([]models.Article, error) {
	res, err := a.session.Query(
		queries.GetAllArticlesQuery,
		nil,
	)
	if err != nil {
		return nil, errors.New("error on reading items from db")
	}

	var articles []models.Article
	for res.Next() {
		var art models.Article
		err = res.Row(&art)
		if err != nil {
			if err == gocb.ErrNoResult {
				return articles, nil
			}
			return nil, err
		}
		articles = append(articles, art)
	}

	return articles, nil

}

func (a *article) GetArticlesBySubCategory(cat, subCat string) ([]models.Article, error) {
	res, err := a.session.Query(
		queries.GetArticlesBySubCategoryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{cat, subCat}},
	)
	if err != nil {
		return nil, errors.New("error on reading items from db")
	}

	var articles []models.Article
	for res.Next() {
		var art models.Article
		err = res.Row(&art)
		if err != nil {
			if err == gocb.ErrNoResult {
				return articles, nil
			}
			return nil, err
		}
		articles = append(articles, art)
	}

	return articles, nil
}

func (a *article) GetArticlesList() ([]string, error) {
	res, err := a.session.Query(
		queries.GetArticlesTitleListQuery,
		nil,
	)
	if err != nil {
		return nil, errors.New("error while quering on db")
	}

	var titles []string
	for res.Next() {
		var title models.ArticleTitle
		err := res.Row(&title)
		if err != nil {
			if err == gocb.ErrNoResult {
				return titles, nil
			}

			return nil, err
		}

		titles = append(titles, title.Title)
	}

	return titles, nil
}
