package repository

import (
	"medikakh/domain/models"

	"github.com/couchbase/gocb/v2"
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
	GetVidoSubCategory(id string) (*string, error)
	GetAllVideos() ([]models.Video, error)
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
	res, err := a.session.Query(
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
	res, err := a.session.Query(
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
