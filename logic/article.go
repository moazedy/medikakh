package logic

import (
	"errors"
	"medikakh/application/utils"
	"medikakh/domain/constants"
	"medikakh/domain/datastore"
	"medikakh/domain/models"
	"medikakh/repository"
	"medikakh/service/authorization"

	"github.com/google/uuid"
)

type ArticleLogic interface {
	SaveArticle(userRole string, art models.Article) error
	ReadArticle(userRole, articleTitle string) (*models.Article, error)
	DeleteArticle(userRole, title string) error
	GetArticleStatus(title string) (*string, error)
	UpdateArticle(userRole string, art models.ArticleUpdate) error
	GetArticlesList(userRole string) ([]string, error)
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

	articleExistance, err := a.repo.IsArticleExists(art.Title)
	if err != nil {
		return err
	}

	if *articleExistance {
		return errors.New("article alredy exists")
	}

	// chacking for existance of determined category in new article
	session, err := datastore.NewCouchbaseSession()
	if err != nil {
		return errors.New("faild to make session with db")
	}
	categoryLogic := NewCategoryLogic(repository.NewCategoryRepo(session))
	err = categoryLogic.IsCategoryExists(art.Category)
	if err != nil {
		return err
	}

	art.Id = uuid.New()
	err = a.repo.Save(art)

	return err
}

func (a *article) ReadArticle(userRole, articleTitle string) (*models.Article, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment invalid")
	}

	status, err := a.GetArticleStatus(articleTitle)
	if err != nil {
		return nil, err
	}

	if *status == "private" {
		ok := authorization.IsPermissioned(userRole, constants.ArticleObject, constants.ReadAction)
		if !ok {
			return nil, errors.New("premission denied")
		}

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

func (a *article) GetArticleStatus(title string) (*string, error) {
	id, err := a.repo.GetArticleId(title)
	if err != nil {
		return nil, err
	}
	status, err := a.repo.GetArticleStatus(*id)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (a *article) UpdateArticle(userRole string, art models.ArticleUpdate) error {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.ArticleObject, constants.UpdateAction)
	if !ok {
		return errors.New("premission denied")
	}

	oldArt, err := a.repo.ReadArticleById(art.Id.String())
	if err != nil {
		return err
	}
	var NewArt models.Article
	NewArt.Id = oldArt.Id
	if art.Title == nil {
		NewArt.Title = oldArt.Title
	} else {
		NewArt.Title = *art.Title
	}

	if art.Status == nil {
		NewArt.Status = oldArt.Status
	} else {
		NewArt.Status = *art.Status
	}

	if art.Summary == nil {
		NewArt.Summary = oldArt.Summary
	} else {
		NewArt.Summary = *art.Summary
	}

	if art.Etiology == nil {
		NewArt.Etiology = oldArt.Etiology
	} else {
		NewArt.Etiology = *art.Etiology
	}

	if art.ClinicalFeatures == nil {
		NewArt.ClinicalFeatures = oldArt.ClinicalFeatures
	} else {
		NewArt.ClinicalFeatures = *art.ClinicalFeatures
	}

	if art.Diagnostics == nil {
		NewArt.Diagnostics = oldArt.Diagnostics
	} else {
		NewArt.Diagnostics = *art.Diagnostics
	}

	if art.Treatment == nil {
		NewArt.Treatment = oldArt.Treatment
	} else {
		NewArt.Treatment = *art.Treatment
	}

	if art.Complications == nil {
		NewArt.Complications = oldArt.Complications
	} else {
		NewArt.Complications = *art.Complications
	}

	if art.Prevention == nil {
		NewArt.Prevention = oldArt.Prevention
	} else {
		NewArt.Prevention = *art.Prevention
	}

	if art.References == nil {
		NewArt.References = oldArt.References
	} else {
		NewArt.References = *art.References
	}

	if art.Category == nil {
		NewArt.Category = oldArt.Category
	} else {
		NewArt.Category = *art.Category
	}

	if art.SubCategory == nil {
		NewArt.SubCategory = oldArt.SubCategory
	} else {
		NewArt.SubCategory = *art.SubCategory
	}

	err = a.repo.UpdateArticle(NewArt)
	if err != nil {
		return err
	}

	return nil

}

func (a *article) GetArticlesList(userRole string) ([]string, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.ArticleObject, constants.ReadAction)
	if !ok {
		return nil, errors.New("premission denied")
	}

	titles, err := a.repo.GetArticlesList()
	if err != nil {
		return nil, err
	}

	return titles, nil
}
