package models

type Movie struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
}
