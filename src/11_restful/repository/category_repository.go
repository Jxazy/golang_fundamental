package repository

import "context"

type CategoryRepository interface {
	Save(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
	FindById(ctx context.Context)
	FindAll(ctx context.Context)
}
