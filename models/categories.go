package models

type Categories struct {
	ID          int    `json`
	Title       string `json:title`
	Description string `json:description`
}
