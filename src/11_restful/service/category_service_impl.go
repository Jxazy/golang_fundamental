package service

import (
	"03_unit_test/repository"
	"11_restful/helper"
	"11_restful/model/web"
	"database/sql"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateReuqest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	//TODO implement me
	panic("implement me")
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	//TODO implement me
	panic("implement me")
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context, categoryId int) web.CategoryResponse {
	//TODO implement me
	panic("implement me")
}

func (service *CategoryServiceImpl) FindById(ctx context.Context) []web.CategoryResponse {
	//TODO implement me
	panic("implement me")
}
