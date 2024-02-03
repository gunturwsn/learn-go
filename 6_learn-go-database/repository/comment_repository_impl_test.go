package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	learn_go_database "learn-go-database"
	"learn-go-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(learn_go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "yaya@gmail.com",
		Comment: "Test Comment Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(learn_go_database.GetConnection())

	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(learn_go_database.GetConnection())

	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
