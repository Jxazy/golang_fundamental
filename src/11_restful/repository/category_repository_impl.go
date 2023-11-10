package repository

import (
	"11_restful/helper"
	"11_restful/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func (repository CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into customer(name) values (?)"

	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.CheckError(err)

	id, err := result.LastInsertId()
	helper.CheckError(err)

	category.Id = int(id)
	return category

}

func (repository CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category SET name = ? where id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.CheckError(err)

	return category
}

func (repository CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.CheckError(err)

}

func (repository CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from category where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.CheckError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&categoryId, &category.Name)
		helper.CheckError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}

func (repository CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.CheckError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.CheckError(err)
		categories = append(categories, category)
	}

	return categories

}
