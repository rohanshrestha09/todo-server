package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/todo/models"
	"github.com/rohanshrestha09/todo/service/database"
)

// Mark Important Todo godoc
//
//	@Summary	Mark important todo
//	@Tags		Todo
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Project ID"
//	@Success	201	{object}	database.Response
//	@Router		/todo/{id}/important [post]
//	@Security	Bearer
func MarkImportant(ctx *gin.Context) {
	todo := ctx.MustGet("todo").(models.Todo)

	data := map[string]any{"Important": true}

	response, err := database.Update(todo, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Todo updated"))

}

// Unmark Important Todo godoc
//
//	@Summary	Unmark important todo
//	@Tags		Todo
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Todo ID"
//	@Success	201	{object}	database.Response
//	@Router		/todo/{id}/important [delete]
//	@Security	Bearer
func UnmarkImportant(ctx *gin.Context) {
	todo := ctx.MustGet("todo").(models.Todo)

	data := map[string]any{"Important": false}

	response, err := database.Update(todo, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Todo updated"))

}
