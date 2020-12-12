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
	ReadArticleFullContent(id string) (*map[string]string, error)
	UpdateArticle(article models.Article) error
	DeleteArticle(id string) error
	GetArticleStatus(id string) (*string, error)
	IsArticleExists(title string) (*bool, error)
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
