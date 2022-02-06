package repository

import "go-unit-test/entity"

type CategoryRepository interface {
	FindById(id int) *entity.Category
}
