package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/todo/models"
	"github.com/rohanshrestha09/todo/service/database"
)

func List() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		authUser := ctx.MustGet("auth").(models.User)

		args := database.GetByIDArgs{}

		data, err := database.GetByID[models.List](ctx.Param("id"), args)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if authUser.ID != data.UserID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		ctx.Set("list", data)

		ctx.Next()

	}
}
