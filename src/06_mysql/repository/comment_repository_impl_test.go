package repository

import (
	mysql "06_mysql"
	"06_mysql/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(mysql.GetConnection())

	ctx := context.Background()
	id := entity.Comment{
		Id: 36,
	}

	result, err := commentRepository.FindById(ctx, id.Id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(mysql.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, result := range result {
		fmt.Println(result)
	}
}

func TestDeleteById(t *testing.T) {
	commentRepository := NewCommentRepository(mysql.GetConnection())

	ctx := context.Background()
	id := entity.Comment{
		Id: 37,
	}

	_, err := commentRepository.DeleteById(ctx, id.Id)
	if err != nil {
		fmt.Println("Delete id", id.Id, "Succesful")
	}

}

func TestUpdateById(t *testing.T) {
	commentRepository := NewCommentRepository(mysql.GetConnection())

	ctx := context.Background()
	id := int32(36)

	comment := entity.Comment{
		Comment: "Test Repository Update END",
	}

	result, err := commentRepository.UpdateById(ctx, id, comment)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Update id", id, "Succesful:", result.Comment)
	}
}
