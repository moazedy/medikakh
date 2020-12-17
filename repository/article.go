package repository

import (
	"medikakh/domain/models"

	"github.com/couchbase/gocb/v2"
)

type ArticleRepo interface {
	Save(article models.Article) error
	ReadArticleById(id string) (*models.Article, error)
	ReadArticleByTitle(title string) (*models.Article, error)
	GetArticleCategory(id string) (*string, error)
	GetArticleSubsCategory(id string) (*string, error)
	GetArticleId(title string) (*string, error)
	ReadArticleSummery(id string) (*string, error)
	ReadArticleResult(id string) (*string, error)
	ReadArticleContent(id string) (*string, error)
	// it shuld go to logic	ReadArticleFullContent(id string) (*map[string]string, error)
	UpdateArticle(article models.Article) error
	DeleteArticle(id string) error
	GetArticleStatus(id string) (*string, error)
	IsArticleExists(title string) (*bool, error)
	GetArticleByCategory(cat string) ([]models.Article, error) // to do
	GetAllArticles() ([]models.Article, error)                 // to do
	GetCategorySubCategories(category string) ([]string, error)
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
		queries,
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
		queries.GetArticleCategory,
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

	return category, nil
}

func (a *article) GetArticleSubsCategory(id string) (*string, error) {
	res, err := a.session.Query(
		queries.GetArticleSubCategory,
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
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
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
func (a *article) ReadArticleSummery(id string) (*string, error) {
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

func (a *article) ReadArticleResult(id string) (*string, error) {
	res, err := a.session.Query(
		queries.ReadArticleResultQuery,
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

func (a *article) ReadArticleContent(id string) (*string, error) {
	res, err := a.session.Query(
		queries.ReadArticleContentQuery,
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
	res, err := u.session.Query(
		queries.IsArticleExistsQuery,
		&gocb.QueryOptions{PositinalParameters: []interface{}{title}},
	)
	if err != nil {
		return nil, errors.New("error on serching for specific article")
	}

	var id string
	for res.Next() {
		err = res.Row(&id)
		if err != nil {
			if err == gocb.ErrNoResult {
				return false, errors.New("article does not exist") //it should be checked later
			}
			return nil, err
		}
	}

	if id != "" || id != nil {
		return true, nil
	}

	return false, nil
}

func (a *article) UpdateArticle(article models.Article) error {
	_, err := a.session.Query(
		queries.UpdateArticleQuery,
		&gocb.QueryOptions{PositionaParameters: []interface{}{
			article.Title,
			article.Status,
			article.Summery,
			article.Content,
			article.Result,
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
