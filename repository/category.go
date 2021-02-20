package repository

import (
	"errors"
	"fmt"
	"log"
	"medikakh/domain/models"
	"medikakh/repository/queries"

	"github.com/couchbase/gocb/v2"
)

type CategoryRepo interface {
	InsertCategory(cat models.Category) error
	ReadCategoryById(id string) (*models.Category, error)
	ReadCategoryByName(name string) (*models.Category, error)
	ReadCategorySubCategories(id string) ([]string, error)
	GetCategoryId(name string) (*string, error)
	GetCategories() ([]models.Category, error)
	IsCategoryExists(name string) error
}

type category struct {
	session *gocb.Cluster
}

func NewCategoryRepo(session *gocb.Cluster) CategoryRepo {
	c := new(category)
	c.session = session

	return c
}

func (c *category) InsertCategory(cat models.Category) error {
	_, err := c.session.Query(
		queries.InsertCategoryQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{cat.Id, cat}},
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *category) ReadCategoryById(id string) (*models.Category, error) {
	res, err := c.session.Query(
		queries.ReadCategoryByIdQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var cat models.Category
	for res.Next() {
		err = res.Row(&cat)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return &cat, nil
}

func (c *category) ReadCategoryByName(name string) (*models.Category, error) {
	res, err := c.session.Query(
		queries.ReadCategoryByNameQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{name}},
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var cat models.Category
	for res.Next() {
		err = res.Row(&cat)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return &cat, nil

}

func (c *category) ReadCategorySubCategories(id string) ([]string, error) {
	res, err := c.session.Query(
		queries.ReadCategorySubCategoriesQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{id}},
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var subCats []string
	for res.Next() {
		err = res.Row(&subCats)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return subCats, nil
}

func (c *category) GetCategoryId(name string) (*string, error) {
	res, err := c.session.Query(
		queries.GetCategoryIdQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{name}},
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var id string
	for res.Next() {
		err = res.Row(&id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return &id, nil
}

func (c *category) GetCategories() ([]models.Category, error) {
	res, err := c.session.Query(
		queries.GetCategoriesQuery,
		nil,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var cats []models.Category
	for res.Next() {
		var cat models.Category
		err = res.Row(&cat)
		if err != nil {
			if err == gocb.ErrNoResult {
				return cats, nil
			}

			log.Println(err)
			return nil, err
		}

		cats = append(cats, cat)
	}

	return cats, nil
}

func (c *category) IsCategoryExists(name string) error {
	res, err := c.session.Query(
		queries.IsCategoryExistsQuery,
		&gocb.QueryOptions{PositionalParameters: []interface{}{name}},
	)
	if err != nil {
		log.Println(err)
		return err
	}

	var count models.Count
	err = res.One(&count)
	if err != nil {
		if err == gocb.ErrNoResult {
			return errors.New("category does not exists")
		}

		log.Println(err)
		return err
	}
	fmt.Println(count.Count)
	if count.Count <= 0 {
		return errors.New("category does not exists")
	}

	return nil
}
