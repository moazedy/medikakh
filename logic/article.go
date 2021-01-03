package logic

import (
	"errors"
	"medikakh/application/utils"
	"medikakh/domain/constants"
	"medikakh/domain/models"
	"medikakh/repository"
	"medikakh/service/authorization"

	"github.com/google/uuid"
)

type ArticleLogic interface {
	SaveArticle(userRole string, art models.Article) error
	ReadArticle(userRole, articleTitle string) (*models.Article, error)
	DeleteArticle(userRole, title string) error
}

type article struct {
	repo repository.ArticleRepo
}

func NewArticleLogic(logic repository.ArticleRepo) ArticleLogic {
	a := new(article)
	a.repo = logic
	return a
}

func (a *article) SaveArticle(userRole string, art models.Article) error {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.ArticleObject, constants.SaveAction)
	if !ok {
		return errors.New("premission denied")
	}

	art.Id = uuid.New()
	err := a.repo.Save(art)

	return err
}

func (a *article) ReadArticle(userRole, articleTitle string) (*models.Article, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment invalid")
	}

	ok := authorization.IsPermissioned(userRole, constants.ArticleObject, constants.ReadAction)
	if !ok {
		return nil, errors.New("premission denied")
	}

	res, err := a.repo.ReadArticleByTitle(articleTitle)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *article) DeleteArticle(userRole, title string) error {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return errors.New("role statment invalid")
	}

	permissionOk := authorization.IsPermissioned(userRole, constants.ArticleObject, constants.DeleteAction)
	if !permissionOk {
		return errors.New("unauthorized user")
	}

	id, err := a.repo.GetArticleId(title)
	if err != nil {
		return err
	}

	err = a.repo.DeleteArticle(*id)
	if err != nil {
		return err
	}

	return nil
}
