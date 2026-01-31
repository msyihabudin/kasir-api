package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]*models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Create(c *models.Category) (*models.Category, error) {
	return s.repo.Create(c)
}

func (s *CategoryService) GetByID(id string) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Update(c *models.Category) (*models.Category, error) {
	return s.repo.Update(c)
}

func (s *CategoryService) Delete(id string) error {
	return s.repo.Delete(id)
}
