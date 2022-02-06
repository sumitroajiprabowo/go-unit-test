package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type CategoryService struct {
	repository repository.CategoryRepository
}

func (service CategoryService) Get(id int) (*entity.Category, error) {
	category := service.repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}
