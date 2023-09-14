package repository

import (
	"06_mysql/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, Comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := repository.DB.ExecContext(ctx, script, Comment.Email, Comment.Comment)
	if err != nil {
		return Comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Comment, err
	}

	Comment.Id = int32(id)
	return Comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	Comment := entity.Comment{}
	if err != nil {
		return Comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&Comment.Id, &Comment.Email, &Comment.Comment)
		return Comment, nil
	} else {
		return Comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}

}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repository *commentRepositoryImpl) DeleteById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "DELETE FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	Comment := entity.Comment{}
	if err != nil {
		return Comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&Comment.Id, &Comment.Email, &Comment.Comment)
		return Comment, nil
	} else {
		return Comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *commentRepositoryImpl) UpdateById(ctx context.Context, id int32, Comment entity.Comment) (entity.Comment, error) {
	script := "UPDATE comments SET comment = ? WHERE id = ? LIMIT 1"
	result, err := repository.DB.ExecContext(ctx, script, Comment.Comment, id)
	if err != nil {
		return Comment, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return Comment, err
	}

	return Comment, nil
}
