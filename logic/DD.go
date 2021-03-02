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

type DDLogic interface {
	InsertData(userRole string, dd models.DDmodel) error
	ReadDD(title string) (*models.DDmodel, error)
	ReadDDbyPattern(pattern string) ([]string, error)
}

type dd struct {
	repo repository.DDrepo
}

func NewDDLogic(repo repository.DDrepo) DDLogic {
	d := new(dd)
	d.repo = repo

	return d
}

func (d *dd) InsertData(userRole string, dd models.DDmodel) error {
	roleCorrectness := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleCorrectness {
		return errors.New("role statment is invalid")
	}

	// checking for user premissins on saving articles
	ok := authorization.IsPermissioned(userRole, constants.DDobject, constants.SaveAction)
	if !ok {
		return errors.New("premission denied")
	}

	if dd.Title == "" || dd.Content == nil {
		return errors.New("title or content, can not be empty")
	}

	ddExistance, err := d.repo.IsDDExists(dd.Title)
	if err != nil {
		return err
	}
	if *ddExistance {
		return errors.New("dd already exists")
	}

	dd.Id = uuid.New()
	err = d.repo.InsertData(dd)
	if err != nil {
		return err
	}

	return nil
}

func (d *dd) ReadDD(title string) (*models.DDmodel, error) {
	if title == "" {
		return nil, errors.New("title can not be empty")
	}

	return d.repo.ReadDataByTitle(title)
}

func (d *dd) ReadDDbyPattern(pattern string) ([]string, error) {
	if pattern == "" {
		return nil, errors.New("title can not be empty")
	}

	dds, err := d.repo.ReadDataUsingPattern(pattern)
	if err != nil {
		return nil, err
	}

	titles := make([]string, len(dds.Titles))
	for i, v := range dds.Titles {
		titles[i] = v.Title
	}

	return titles, nil
}
