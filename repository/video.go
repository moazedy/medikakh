package repository

import (
	"errors"
	"github.com/couchbase/gocb/v2"
	"medikakh/domain/models"
	"medikakh/repository/queries"
)

type VideoRepo interface {
	Save(video models.Video) error
	GetVideoById(id string) (*models.Video, error)
	GetVideoByTitle(title string) (*models.Video, error)
	GetVideoId(title string) (*string, error)
	GetVideoCategory(id string) (*string, error)
	GetVideoSubCategory(id string) (*string, error)
	GetVideosByCategory(cat string) ([]models.Video, error)
	DeleteVideo(id string) error
	UpdateVideo(video models.Video) error
	GetVideoStatus(id string) (*string, error)
	GetAllVideos() ([]models.Video, error)
	GetVideoBySubCategory(cat, subCat string) ([]models.Video, error)
	IsVideoExists(title string) (*bool, error)
}

type video struct {
	session *gocb.Cluster
}

func NewVideoRepo(session *gocb.Cluster) VideoRepo {
	v := new(video)
	v.session = session
	return v
}

func (v *video) Save(video models.Video) error {
	_, err := v.session.Query(
		queries.SaveVideoQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"id":    video.Id,
			"video": video,
		}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (v *video) GetVideoById(id string) (*models.Video, error) {
	res, err := v.session.Query(
		queries.GetVideoByIdQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"id": id,
		}},
	)
	if err != nil {
		return nil, err
	}

	var newVideo models.Video
	for res.Next() {
		err = res.Row(&newVideo)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("the video does not exist !")
			}

			return nil, err
		}
	}

	return &newVideo, nil

}

func (v *video) GetVideoByTitle(title string) (*models.Video, error) {
	res, err := v.session.Query(
		queries.GetVideoByTitleQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"title": title,
		}},
	)
	if err != nil {
		return nil, err
	}

	var newVideo models.Video
	for res.Next() {
		err = res.Row(&newVideo)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("the video does not exist !")
			}

			return nil, err
		}
	}

	return &newVideo, nil

}

func (v *video) GetVideoId(title string) (*string, error) {
	res, err := v.session.Query(
		queries.GetVideoIdQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{title}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("video does not exist !")
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

func (v *video) GetVideoCategory(id string) (*string, error) {
	res, err := v.session.Query(
		queries.GetVideoCategoryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("video does not exist !")
		}

		return nil, err
	}

	var cat string
	for res.Next() {
		err = res.Row(&cat)
		if err != nil {
			return nil, err
		}
	}

	return &cat, nil

}

func (v *video) GetVideoSubCategory(id string) (*string, error) {
	res, err := v.session.Query(
		queries.GetVideoSubCategoryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("video does not exist !")
		}

		return nil, err
	}

	var subCat string
	for res.Next() {
		err = res.Row(&subCat)
		if err != nil {
			return nil, err
		}
	}

	return &subCat, nil

}
func (v *video) GetVideosByCategory(cat string) ([]models.Video, error) {
	res, err := v.session.Query(
		queries.GetVideosByCategoryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{cat}},
	)
	if err != nil {
		return nil, err
	}

	var videos []models.Video
	for res.Next() {
		var video models.Video
		err = res.Row(&video)
		if err != nil {
			if err == gocb.ErrNoResult {
				return videos, nil
			}

			return nil, err
		}

		videos = append(videos, video)
	}

	return videos, nil
}

func (v *video) DeleteVideo(id string) error {
	_, err := v.session.Query(
		queries.DeleteVideoQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (v *video) UpdateVideo(video models.Video) error {
	_, err := v.session.Query(
		queries.UpdateVideoQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"title":        video.Title,
			"contentlink":  video.ContentLink,
			"status":       video.Status,
			"category":     video.Category,
			"subcategory":  video.SubCategory,
			"descriptions": video.Descriptions,
		}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (v *video) GetVideoStatus(id string) (*string, error) {
	res, err := v.session.Query(
		queries.GetVideoStatusQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("video does not exist !")
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

func (v *video) GetAllVideos() ([]models.Video, error) {
	res, err := v.session.Query(
		queries.GetAllVideosQuery,
		nil,
	)
	if err != nil {
		return nil, err
	}

	var videos []models.Video
	for res.Next() {
		var video models.Video
		err = res.Row(&video)
		if err != nil {
			if err == gocb.ErrNoResult {
				return videos, nil
			}
			return nil, err
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func (v *video) GetVideoBySubCategory(cat, subCat string) ([]models.Video, error) {
	res, err := v.session.Query(
		queries.GetVideoBySubCategoryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{cat, subCat}},
	)
	if err != nil {
		return nil, err
	}

	var videos []models.Video
	for res.Next() {
		var video models.Video
		err = res.Row(&video)
		if err != nil {
			if err == gocb.ErrNoResult {
				return videos, nil
			}

			return nil, err
		}

		videos = append(videos, video)
	}

	return videos, nil

}

func (v *video) IsVideoExists(title string) (*bool, error) {
	res, err := v.session.Query(
		queries.IsVideoExistsQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{title}},
	)
	if err != nil {
		return nil, errors.New("error on serching for specific video")
	}
	var returnValue bool
	var id string
	for res.Next() {
		err = res.Row(&id)
		if err != nil {
			if err == gocb.ErrNoResult {
				return &returnValue, nil
			}
			return nil, err
		}
	}

	if id != "" {
		returnValue = true
	}

	return &returnValue, nil
}
