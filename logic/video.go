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
	Update(userRole string, vid models.VideoUpdate) error
	GetVideosByCategory(userRole, cat string) ([]string, error)
	GetAllVideosList(userRole string) ([]string, error)
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

func (v *video) Update(userRole string, vid models.VideoUpdate) error {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return errors.New("role statment not valid")
	}

	permissionOk := authorization.IsPermissioned(userRole, constants.VideoObject, constants.UpdateAction)
	if !permissionOk {
		return errors.New("unauthorized user")
	}
	// this part seems to be extra. when you reading a file all the parts will be done auto ...
	title, err := v.repo.GetVideoTitle(vid.Id.String())
	if err != nil {
		return err
	}
	videoExistance, err := v.repo.IsVideoExists(*title)
	if err != nil {
		return err
	}
	if !*videoExistance {
		return errors.New("video does not exists")
	}
	// *********************** thill here *********************

	oldVid, err := v.repo.GetVideoById(vid.Id.String())
	if err != nil {
		return err
	}
	var newVid models.Video
	newVid.Id = vid.Id

	if vid.Title != nil {
		newVid.Title = *vid.Title
	} else {
		newVid.Title = oldVid.Title
	}

	if vid.ContentLink != nil {
		newVid.ContentLink = *vid.ContentLink
	} else {
		newVid.ContentLink = oldVid.ContentLink
	}

	if vid.Status != nil {
		newVid.Status = *vid.Status
	} else {
		newVid.Status = oldVid.Status
	}

	if vid.Category != nil {
		// checking for category value validation
		dbSession, err := datastore.NewCouchbaseSession()
		if err != nil {
			return err
		}
		categoryLogic := NewCategoryLogic(repository.NewCategoryRepo(dbSession))

		err = categoryLogic.IsCategoryExists(*vid.Category)
		if err != nil {
			return err
		}
		newVid.Category = *vid.Category
	} else {
		newVid.Category = oldVid.Category
	}

	if vid.SubCategory != nil {
		newVid.SubCategory = *vid.SubCategory
	} else {
		newVid.SubCategory = oldVid.SubCategory
	}

	if vid.Descriptions != nil {
		newVid.Descriptions = *vid.Descriptions
	} else {
		newVid.Descriptions = oldVid.Descriptions
	}

	err = v.repo.UpdateVideo(newVid)
	if err != nil {
		return err
	}

	return nil
}

func (v *video) GetVideosByCategory(userRole, cat string) ([]string, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment not valid")
	}

	permissionOk := authorization.IsPermissioned(userRole, constants.VideoObject, constants.ReadAction)
	if !permissionOk {
		return nil, errors.New("unauthorized user")
	}

	// checking for category value validation
	dbSession, err := datastore.NewCouchbaseSession()
	if err != nil {
		return nil, err
	}
	categoryLogic := NewCategoryLogic(repository.NewCategoryRepo(dbSession))

	err = categoryLogic.IsCategoryExists(cat)
	if err != nil {
		return nil, err
	}

	vids, err := v.repo.GetVideosByCategory(cat)
	if err != nil {
		return nil, err
	}

	return vids, nil
}

func (v *video) GetAllVideosList(userRole string) ([]string, error) {
	roleOK := utils.CheckForRoleStatmentCorrectness(userRole)
	if !roleOK {
		return nil, errors.New("role statment not valid")
	}

	permissionOk := authorization.IsPermissioned(userRole, constants.VideoObject, constants.ReadAction)
	if !permissionOk {
		return nil, errors.New("unauthorized user")
	}

	vids, err := v.repo.GetAllVideosList()
	if err != nil {
		return nil, err
	}

	return vids, nil
}
