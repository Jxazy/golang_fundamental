package service

import (
	"11_restful/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateReuqest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindAll(ctx context.Context, categoryId int) web.CategoryResponse
	FindById(ctx context.Context) []web.CategoryResponse
}
