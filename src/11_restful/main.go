package main

import (
	"06_mysql/repository"
	"11_restful/app"
	"11_restful/controller"
	"11_restful/service"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.New
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("api/categories/:categoryId", categoryController.FindById)
	router.POST("api/categories", categoryController.Create)
	router.PUT(("api/categories/:categoryId", categoryController.Update)
	router.DELETE("api/categories/:categoryId", categoryController.Delete)
}
