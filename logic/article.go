package logic

import (
	"errors"
	"medikakh/domain/constants"
	"medikakh/domain/models"
	"medikakh/repository"
	"medikakh/service/authorization"

	"github.com/google/uuid"
)

type ArticleLogic interface {
	SaveArticle(userRole string, art models.Article) error
	ReadArticle(userRole, articleTitle string) (*models.Article, error)
}

type article struct {
	repo repository.ArticleRepo
}

func NewArticleLogic(logic repository.ArticleRepo) ArticleLogic {
	a := new(article)
	a.logic = logic
	return a
}

func (a *article) SaveArticle(userRole string, art models.Article) error {
	// checking for user premissins on saving articles
	ok := authorization.IsPremissioned(userRole, constants.ArticleObject, constants.SaveAction)
	if !ok {
		return errors.New("premission denied")
	}

	art.Id = uuid.New()
	err := a.repo.Save(art)

	return err
}

func (a *article) ReadArticle(userRole, articleTitle string) (*models.Article, error) {
	ok := authorization.IsPremissioned(userRole, constants.ArticleObject, constants.ReadAction)
	if !ok {
		return nil, errors.New("premission denied")
	}

	res, err := a.repo.ReadArticleByTitle(articleTitle)
	if err != nil {
		return nil, err
	}

	return res, nil
}
