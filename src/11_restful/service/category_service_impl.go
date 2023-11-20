package service

import (
	"11_restful/helper"
	"11_restful/model/domain"
	"11_restful/model/web"
	"11_restful/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validator *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validator:          validator,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateReuqest) web.CategoryResponse {
	err := service.Validator.Struct(request)
	helper.CheckError(err)

	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.CategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validator.Struct(request)
	helper.CheckError(err)

	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.CheckError(err)

	category.Name = request.Name

	service.CategoryRepository.Update(ctx, tx, category)

	return helper.CategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.CheckError(err)

	service.CategoryRepository.Delete(ctx, tx, category)

}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.CheckError(err)

	return helper.CategoryResponse(category)

}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)

}
