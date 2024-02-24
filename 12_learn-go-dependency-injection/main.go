package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"learn-go-dependency-injection/app"
	"learn-go-dependency-injection/controller"
	"learn-go-dependency-injection/helper"
	"learn-go-dependency-injection/middleware"
	"learn-go-dependency-injection/repository"
	"learn-go-dependency-injection/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
