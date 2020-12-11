package models

type Video struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	ContentLink  string `json:"content_link"`
	Status       string `json:"status"`
	Category     string `json:"category"`
	SubCategory  string `json:"sub_category"`
	Descriptions string `json:"descriptions"`
}
