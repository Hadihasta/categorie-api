package repositories

import (
	"categories-api/models"
	"database/sql"
	"errors"
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

func (repo *CategoryRepository) GetByID(id int) (*models.Categories, error) {
	query := `
		SELECT id, title, description
		FROM categories
		WHERE id = $1
	`

	var c models.Categories
	err := repo.db.QueryRow(query, id).Scan(&c.ID, &c.Title, &c.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("category not found")
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (repo *CategoryRepository) Create(category *models.Categories) error {
	query := "INSERT INTO Categories (title, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, category.Title, category.Description).Scan(&category.ID)
	return err
}


func (repo *CategoryRepository) Update(category *models.Categories) error {
	query := `
		UPDATE categories
		SET title = $1, description = $2
		WHERE id = $3
	`

	result, err := repo.db.Exec(
		query,
		category.Title,
		category.Description,
		category.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("category not found")
	}

	return nil
}


func (repo *CategoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"

	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("category not found")
	}

	return nil
}