package repository

import (
	"06_mysql/entity"
	"context"
)

type CommentRepository interface {
	Insert(ctx context.Context, Comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
	DeleteById(ctx context.Context, id int32) (entity.Comment, error)
	UpdateById(ctx context.Context, id int32, Comment entity.Comment) (entity.Comment, error)
}
