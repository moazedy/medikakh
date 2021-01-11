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

type ArticleUpdate struct {
	Id               uuid.UUID `json:"id"`
	Title            *string   `json:"title,omitempty"`
	Status           *string   `json:"status,omitempty"`
	Summary          *string   `json:"summary,omitempty"`
	Etiology         *string   `json:"etiology,omitempty"`
	ClinicalFeatures *string   `json:"clinical_features,omitempty"`
	Diagnostics      *string   `json:"diagnostics,omitempty"`
	Treatment        *string   `json:"treatment,omitempty"`
	Complications    *string   `json:"complications,omitempty"`
	Prevention       *string   `json:"prevention,omitempty"`
	References       *string   `json:"references,omitempty"`
	Category         *string   `json:"category,omitempty"`
	SubCategory      *string   `json:"sub_category,omitempty"`
}

type ArticleTitle struct {
	Title string `json:"title"`
}
