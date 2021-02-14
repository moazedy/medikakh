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

type CategoryLogic interface {
	InsertCategory(userRole string, cat models.Category) error
	GetCategory(userRole, catTitle string) (*models.Category, error)
	GetCategories(userRole string) ([]models.Category, error)
	GetCategorySubCategories(userRole, catTitle string) ([]string, error)
	IsCategoryExists(name string) error
}

type category struct {
	repo repository.CategoryRepo
}

func NewCategoryLogic(repo repository.CategoryRepo) CategoryLogic {
	c := new(category)
	c.repo = repo

	return c
}

func (c *category) InsertCategory(userRole string, cat models.Category) error {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.CategoryObject, constants.SaveAction)
	if !ok {
		return errors.New("premission denied")
	}

	err := c.repo.IsCategoryExists(cat.Name)
	if err == nil {
		return errors.New("category alredy exists")
	}

	cat.Id = uuid.New()
	err = c.repo.InsertCategory(cat)
	if err != nil {
		return err
	}

	return nil
}

func (c *category) GetCategory(userRole, catTitle string) (*models.Category, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.CategoryObject, constants.ReadAction)
	if !ok {
		return nil, errors.New("premission denied")
	}

	gotCategory, err := c.repo.ReadCategoryByName(catTitle)
	if err != nil {
		return nil, err
	}

	return gotCategory, nil

}

func (c *category) GetCategories(userRole string) ([]models.Category, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.CategoryObject, constants.ReadAction)
	if !ok {
		return nil, errors.New("premission denied")
	}

	categories, err := c.repo.GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *category) GetCategorySubCategories(userRole, catTitle string) ([]string, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.CategoryObject, constants.ReadAction)
	if !ok {
		return nil, errors.New("premission denied")
	}

	categoryId, err := c.repo.GetCategoryId(catTitle)
	if err != nil {
		return nil, err
	}
	subs, err := c.repo.ReadCategorySubCategories(*categoryId)
	if err != nil {
		return nil, err
	}

	return subs, nil
}

func (c *category) IsCategoryExists(name string) error {
	if name == "" {
		return errors.New("category name can not be empty")
	}

	err := c.repo.IsCategoryExists(name)
	if err != nil {
		return err
	}

	return nil
}
