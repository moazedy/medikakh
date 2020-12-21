package models

import "github.com/google/uuid"

type Category struct {

	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	SubCategories []string  `json:"sub_categories"`
}
