package models

import "github.com/google/uuid"

type DDmodel struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content []string  `json:"content"`
}

type DDtitle struct {
	Title string `josn:"title"`
}

type DDtitles struct {
	Titles []DDtitle `json:"titles"`
}
