package services

import (
	"categories-api/models"
	"categories-api/repositories"
	"errors"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (c *CategoryService) GetAll() ([]models.Categories, error) {
	return c.repo.GetAll()
}

func (s *CategoryService) GetByID(id int) (*models.Categories, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Create(category *models.Categories) error {
	// jika title kurang dari 3 characther
	if len(category.Title) < 3 {
		return errors.New("title too short")
	}

	return s.repo.Create(category)
}

func (s *CategoryService) Update(category *models.Categories) error {
	if category.ID == 0 {
		return errors.New("invalid category id")
	}
	if len(category.Title) < 3 {
		return errors.New("title must be at least 3 characters")
	}
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
