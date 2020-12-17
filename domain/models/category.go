package models

type Category struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	SubCategories []string  `json:"sub_categories"`
}
