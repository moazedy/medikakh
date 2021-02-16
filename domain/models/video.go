package models

import "github.com/google/uuid"

type Video struct {
	Id           uuid.UUID `json:"id,omitempty"`
	Title        string    `json:"title"`
	ContentLink  string    `json:"content_link"`
	Status       string    `json:"status"`
	Category     string    `json:"category"`
	SubCategory  string    `json:"sub_category"`
	Descriptions string    `json:"descriptions"`
}

type VideoUpdate struct {
	Id           uuid.UUID `json:"id"`
	Title        *string   `json:"title,omitempty"`
	ContentLink  *string   `json:"content_link,omitempty"`
	Status       *string   `json:"status,omitempty"`
	Category     *string   `json:"category,omitempty"`
	SubCategory  *string   `json:"sub_category,omitempty"`
	Descriptions *string   `json:"descriptions,omitempty"`
}

type VideoTitle struct {
	Title string `json:"title"`
}
