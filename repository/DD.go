package repository

import (
	"medikakh/domain/models"
	"medikakh/repository/queries"

	"github.com/couchbase/gocb/v2"
)

type DDrepo interface {
	InsertData(dd models.DDmodel) error
	ReadDataById(Id string) (*models.DDmodel, error)
	ReadDataByTitle(title string) (*models.DDmodel, error)
	ReadDataUsingPattern(title string) ([]models.DDmodel, error)
}

type dd struct {
	session *gocb.Cluster
}

func NewDDrepo(session *gocb.Cluster) DDrepo {
	d := new(dd)
	d.session = session
	return d
}

func (d *dd) InsertData(dd models.DDmodel) error {
	_, err := d.session.Query(
		queries.InsertDataQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"key":   dd.Id,
			"value": dd,
		}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *dd) ReadDataById(Id string) (*models.DDmodel, error) {
	res, err := d.session.Query(
		queries.ReadDataByIdQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{Id}},
	)
	if err != nil {
		return nil, err
	}

	var dd models.DDmodel
	for res.Next() {
		err = res.Row(&dd)
		if err != nil {
			return nil, err
		}
	}

	return &dd, nil
}

func (d *dd) ReadDataByTitle(title string) (*models.DDmodel, error) {
	res, err := d.session.Query(
		queries.ReadDataByTitleQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{title}},
	)
	if err != nil {
		return nil, err
	}

	var dd models.DDmodel
	for res.Next() {
		err = res.Row(&dd)
		if err != nil {
			return nil, err
		}
	}

	return &dd, nil
}
func (d *dd) ReadDataUsingPattern(title string) ([]models.DDmodel, error) {
	pattern := title + "_%"
	res, err := d.session.Query(
		queries.ReadDataUsingPatternQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{pattern}},
	)
	if err != nil {
		return nil, err
	}

	var dds []models.DDmodel
	for res.Next() {
		var dd models.DDmodel
		err = res.Row(&dd)
		if err != nil {
			if err == gocb.ErrNoResult {
				return dds, nil
			}
			return nil, err
		}

		dds = append(dds, dd)
	}

	return dds, nil
}
