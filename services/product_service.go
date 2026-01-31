package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]*models.ProductWithCategory, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(p *models.Product) (*models.Product, error) {
	return s.repo.Create(p)
}

func (s *ProductService) GetByID(id string) (*models.ProductWithCategory, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(p *models.Product) (*models.Product, error) {
	return s.repo.Update(p)
}

func (s *ProductService) Delete(id string) error {
	return s.repo.Delete(id)
}
