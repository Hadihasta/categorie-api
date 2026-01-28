package services

import (
	"categories-api/models"
	"categories-api/repositories"
	"errors"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService{
	return  &CategoryService{repo : repo}
}


func (c *CategoryService) GetAll() ([]models.Categories, error){
	return c.repo.GetAll()
}


func (s *CategoryService) Create(category *models.Categories) error {
	// jika title kurang dari 3 characther
	if len(category.Title) < 3 {
		return errors.New("title too short")
	}

	return s.repo.Create(category)
}


