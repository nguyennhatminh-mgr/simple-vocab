package models

type Vocabulary struct {
	Id    int   `json:"id"`
	Word  string `json:"word"`
	Meaning string `json:"meaning"`
	CategoryId int `json:"category_id"`
	CreatedAt string `json:"created_at"`
}