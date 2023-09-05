package repository

import "03_unit_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
