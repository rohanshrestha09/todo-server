package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/todo/controllers/list"
	"github.com/rohanshrestha09/todo/middleware"
)

func listRouter(router *gin.RouterGroup) {

	router.Use(middleware.Auth())

	router.GET("/", list.GetAll)

	router.POST("/", list.Create)

	router.Use(middleware.List())
	{
		router.GET("/:id", list.Get)

		router.PATCH("/:id", list.Update)

		router.DELETE("/:id", list.Delete)

		router.GET("/:id/todo", list.GetTodos)
	}

}
