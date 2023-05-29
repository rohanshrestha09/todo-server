package routes

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.RouterGroup) {

	authRouter(router.Group("/auth"))

	userRouter(router.Group("/user"))

	ssoRouter(router.Group("/sso"))

	listRouter(router.Group("/list"))

	todoRouter(router.Group("/todo"))

}
