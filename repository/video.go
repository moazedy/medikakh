package repository

import (
	"errors"
	"log"
	"medikakh/domain/models"
	"medikakh/repository/queries"

	"github.com/couchbase/gocb/v2"
)

type VideoRepo interface {
	Save(video models.Video) error
	GetVideoById(id string) (*models.Video, error)
	GetVideoByTitle(title string) (*models.Video, error)
	GetVideoId(title string) (*string, error)
	GetVideoCategory(id string) (*string, error)
	GetVideoSubCategory(id string) (*string, error)
	GetVideosByCategory(cat string) ([]string, error)
	DeleteVideo(id string) error
	UpdateVideo(vid models.Video) error
	GetVideoStatus(id string) (*string, error)
	GetAllVideosList() ([]string, error)
	GetVideoBySubCategory(cat, subCat string) ([]string, error)
	IsVideoExists(title string) (*bool, error)
	GetVideoTitle(id string) (*string, error)
}

type video struct {
	session *gocb.Cluster
}

func NewVideoRepo(session *gocb.Cluster) VideoRepo {
	v := new(video)
	v.session = session
	return v
}

func (v *video) Save(vid models.Video) error {
	_, err := v.session.Query(
		queries.SaveVideoQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"id":    vid.Id,
			"video": vid,
		}},
	)
	if err != nil {
		log.Println(err.Error())
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
		log.Println(err.Error())
		return nil, err
	}

	var newVideo models.Video
	for res.Next() {
		err = res.Row(&newVideo)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("the video does not exist !")
			}

			log.Println(err.Error())
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
		log.Println(err.Error())
		return nil, err
	}

	var newVideo models.Video
	for res.Next() {
		err = res.Row(&newVideo)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("the video does not exist !")
			}

			log.Println(err.Error())
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

		log.Println(err.Error())
		return nil, err
	}

	var id models.Id
	for res.Next() {
		err = res.Row(&id)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	return &id.Id, nil
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

		log.Println(err.Error())
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

		log.Println(err.Error())
		return nil, err
	}

	var subCat string
	for res.Next() {
		err = res.Row(&subCat)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	return &subCat, nil

}
func (v *video) GetVideosByCategory(cat string) ([]string, error) {
	res, err := v.session.Query(
		queries.GetVideosByCategoryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{cat}},
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var videos []string
	for res.Next() {
		var video models.VideoTitle
		err = res.Row(&video)
		if err != nil {
			if err == gocb.ErrNoResult {
				return videos, nil
			}

			log.Println(err.Error())
			return nil, err
		}

		videos = append(videos, video.Title)
	}

	return videos, nil
}

func (v *video) DeleteVideo(id string) error {
	_, err := v.session.Query(
		queries.DeleteVideoQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (v *video) UpdateVideo(vid models.Video) error {
	_, err := v.session.Query(
		queries.UpdateVideoQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{vid.Title,
			vid.ContentLink,
			vid.Status,
			vid.Category,
			vid.SubCategory,
			vid.Descriptions,
			vid.Id}},
	)
	if err != nil {
		log.Println(err.Error())
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

		log.Println(err.Error())
		return nil, err
	}

	var status string
	for res.Next() {
		err = res.Row(&status)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	return &status, nil

}

func (v *video) GetAllVideosList() ([]string, error) {
	res, err := v.session.Query(
		queries.GetAllVideosQuery,
		nil,
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var videos []string
	for res.Next() {
		var video models.VideoTitle
		err = res.Row(&video)
		if err != nil {
			if err == gocb.ErrNoResult {
				return videos, nil
			}
			log.Println(err.Error())
			return nil, err
		}
		videos = append(videos, video.Title)
	}

	return videos, nil
}

func (v *video) GetVideoBySubCategory(cat, subCat string) ([]string, error) {
	res, err := v.session.Query(
		queries.GetVideoBySubCategoryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{cat, subCat}},
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var videos []string
	for res.Next() {
		var video models.VideoTitle
		err = res.Row(&video)
		if err != nil {
			if err == gocb.ErrNoResult {
				return videos, nil
			}

			log.Println(err.Error())
			return nil, err
		}

		videos = append(videos, video.Title)
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
	var id models.Id
	for res.Next() {
		err = res.Row(&id)
		if err != nil {
			if err == gocb.ErrNoResult {
				return &returnValue, nil
			}
			log.Println(err.Error())
			return nil, err
		}
	}

	if id.Id != "" {
		returnValue = true
	}

	return &returnValue, nil
}

func (v *video) GetVideoTitle(id string) (*string, error) {
	res, err := v.session.Query(
		queries.GetVideoTitleQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, errors.New("video does not exist !")
		}

		log.Println(err.Error())
		return nil, err
	}

	var title models.Tilte
	for res.Next() {
		err = res.Row(&title)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	return &title.Title, nil

}
