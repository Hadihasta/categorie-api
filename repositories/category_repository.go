package repositories

import (
	"categories-api/models"
	"database/sql"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAll() ([]models.Categories, error) {
	query := `
		SELECT id, title, description
		FROM categories
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	category := make([]models.Categories, 0)
	for rows.Next() {
		var c models.Categories
		err := rows.Scan(&c.ID, &c.Title, &c.Description)
		if err != nil {
			return nil, err
		}
		category = append(category, c)
	}

	return category, nil
}



func (repo *CategoryRepository) Create(category *models.Categories) error {
	query := "INSERT INTO Categories (title, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, category.Title, category.Description).Scan(&category.ID)
	return err
}
