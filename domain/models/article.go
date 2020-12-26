package models

import "github.com/google/uuid"

type Article struct {
	Id               uuid.UUID `json:"id"`
	Title            string    `json:"title"`
	Status           string    `json:"status"`
	Summary          string    `json:"summary"`
	Etiology         string    `json:"etiology"`
	ClinicalFeatures string    `json:"clinical_features"`
	Diagnostics      string    `json:"diagnostics"`
	Treatment        string    `json:"treatment"`
	Complications    string    `json:"complications"`
	Prevention       string    `json:"prevention"`
	References       string    `json:"references"`
	Category         string    `json:"category"`
	SubCategory      string    `json:"sub_category"`
}
