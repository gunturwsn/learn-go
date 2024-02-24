package app

import (
	"github.com/julienschmidt/httprouter"
	"learn-go-dependency-injection/controller"
	"learn-go-dependency-injection/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
