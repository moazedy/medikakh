package models

type Article struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Summery     string `json:"summery"`
	Content     string `json:"content"`
	Result      string `json:"result"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
}
