package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/todo/controllers/user"
)

func userRouter(router *gin.RouterGroup) {

	router.POST("/register", user.Register)

	router.POST("/login", user.Login)

}
