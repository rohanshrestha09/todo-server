package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/rohanshrestha09/todo/controllers/todo"
	"github.com/rohanshrestha09/todo/middleware"
)

func todoRouter(router *gin.RouterGroup) {

	router.Use(middleware.Auth())

	router.GET("/", todo.GetAll)

	router.POST("/:id", middleware.List(), todo.Create)

	router.Use(middleware.Todo())
	{
		router.GET("/:id", todo.Get)

		router.PATCH("/:id", todo.Update)

		router.DELETE("/:id", todo.Delete)

		router.PUT("/:id/important", todo.MarkImportant)

		router.DELETE("/:id/important", todo.UnmarkImportant)

		router.PUT("/:id/complete", todo.MarkComplete)

		router.DELETE("/:id/complete", todo.UnmarkComplete)
	}

}
