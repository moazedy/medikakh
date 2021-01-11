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

type VideoLogic interface {
	Save(userRole string, vi models.Video) error
	GetVideo(userRole, vidTitle string) (*models.Video, error)
	Delete(userRole, vidTitle string) error
}

type video struct {
	repo repository.VideoRepo
}

func NewVideoLogic(repo repository.VideoRepo) VideoLogic {
	v := new(video)
	v.repo = repo

	return v
}

func (v *video) Save(userRole string, vi models.Video) error {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return errors.New("role statment not valid")
	}

	permissionOk := authorization.IsPermissioned(userRole, constants.VideoObject, constants.SaveAction)
	if !permissionOk {
		return errors.New("unauthorized user")
	}

	videoExistance, err := v.repo.IsVideoExists(vi.Title)
	if err != nil {
		return err
	}
	if *videoExistance {
		return errors.New("video alredy exists")
	}

	// chacking for existance of determined category in new article
	session, err := datastore.NewCouchbaseSession()
	if err != nil {
		return errors.New("faild to make session with db")
	}
	categoryLogic := NewCategoryLogic(repository.NewCategoryRepo(session))
	err = categoryLogic.IsCategoryExists(vi.Category)
	if err != nil {
		return err
	}

	vi.Id = uuid.New()
	err = v.repo.Save(vi)
	if err != nil {
		return err
	}

	return nil
}
func (v *video) GetVideo(userRole, vidTitle string) (*models.Video, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment not valid")
	}

	permissionOk := authorization.IsPermissioned(userRole, constants.VideoObject, constants.ReadAction)
	if !permissionOk {
		return nil, errors.New("unauthorized user")
	}

	videoExistance, err := v.repo.IsVideoExists(vidTitle)
	if err != nil {
		return nil, err
	}
	if !*videoExistance {
		return nil, errors.New("video does not exists")
	}

	wantedVideo, err := v.repo.GetVideoByTitle(vidTitle)
	if err != nil {
		return nil, err
	}

	return wantedVideo, nil

}

func (v *video) Delete(userRole, vidTitle string) error {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return errors.New("role statment not valid")
	}

	permissionOk := authorization.IsPermissioned(userRole, constants.VideoObject, constants.DeleteAction)
	if !permissionOk {
		return errors.New("unauthorized user")
	}

	id, err := v.repo.GetVideoId(vidTitle)
	if err != nil {
		return err
	}
	err = v.repo.DeleteVideo(*id)
	if err != nil {
		return err
	}

	return nil
}
