package models

import "github.com/google/uuid"

type DDmodel struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content []string  `json:"content"`
}
