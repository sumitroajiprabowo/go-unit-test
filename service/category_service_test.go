package service

import (
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{categoryRepository}

func TestCategoryService_GetFail(t *testing.T) {

	categoryRepository.Mock.On("FindById", 1).Return(nil)

	result, err := categoryService.Get(1)
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestCategoryService_GetSuccess(t *testing.T) {

	categoryRepository.Mock.On("FindById", 2).Return(&entity.Category{Id: 1, Name: "test"})

	category, err := categoryService.Get(2)
	assert.Equal(t, 1, category.Id)
	assert.Equal(t, "test", category.Name)
	assert.Nil(t, err)
	assert.NotNil(t, category)

}
