package models

type DDmodel struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Content []string `json:"content"`
}
