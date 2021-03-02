package repository

import (
	"errors"
	"log"
	"medikakh/domain/models"
	"medikakh/repository/queries"

	"github.com/couchbase/gocb/v2"
)

type DDrepo interface {
	InsertData(dd models.DDmodel) error
	ReadDataById(Id string) (*models.DDmodel, error)
	ReadDataByTitle(title string) (*models.DDmodel, error)
	ReadDataUsingPattern(title string) (*models.DDtitles, error)
	IsDDExists(title string) (*bool, error)
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
		log.Println("error is here")
		log.Println(err.Error())
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
		log.Println(err.Error())
		return nil, err
	}

	var dd models.DDmodel
	for res.Next() {
		err = res.Row(&dd)
		if err != nil {
			log.Println(err.Error())
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
		log.Println(err.Error())
		return nil, err
	}

	var dd models.DDmodel
	for res.Next() {
		err = res.Row(&dd)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("dd does not exists")
			}
			log.Println(err.Error())
			return nil, err
		}
	}

	return &dd, nil
}
func (d *dd) ReadDataUsingPattern(title string) (*models.DDtitles, error) {
	pattern := title + "_%"
	res, err := d.session.Query(
		queries.ReadDataUsingPatternQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{pattern}},
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var dds models.DDtitles
	for res.Next() {
		var dd models.DDtitle
		err = res.Row(&dd)
		if err != nil {
			if err == gocb.ErrNoResult {
				return &dds, nil
			}
			log.Println(err.Error())
			return nil, err
		}

		dds.Titles = append(dds.Titles, dd)
	}

	return &dds, nil
}

func (d *dd) IsDDExists(title string) (*bool, error) {
	res, err := d.session.Query(
		queries.IsDDExistsQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{title}},
	)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error on serching for specific dd")
	}
	var returnValue bool
	var id models.Id
	for res.Next() {
		err = res.Row(&id)
		if err != nil {
			if err == gocb.ErrNoResult {
				return &returnValue, errors.New("dd does not exist") //it should be checked later
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
