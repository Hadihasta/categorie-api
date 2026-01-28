package models

type Product struct {
	ID    int    `json`
	Name  string `json:"name`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}