package services

import (
	"categories-api/models"
	"categories-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]models.Product, error){
	return s.repo.GetAll()
}


func (s *ProductService) Create(data *models.Product) error {
	return s.repo.Create(data)
}