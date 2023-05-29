package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/todo/enums"
	"github.com/rohanshrestha09/todo/models"
	"github.com/rohanshrestha09/todo/service/database"
)

// Mark Complete Todo godoc
//
//	@Summary	Mark complete todo
//	@Tags		Todo
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Todo ID"
//	@Success	201	{object}	database.Response
//	@Router		/todo/{id}/complete [post]
//	@Security	Bearer
func MarkComplete(ctx *gin.Context) {
	todo := ctx.MustGet("todo").(models.Todo)

	data := models.Todo{
		Status: enums.COMPLETED,
	}

	response, err := database.Update(todo, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Todo updated"))

}

// Unmark Complete Todo godoc
//
//	@Summary	Unmark complete todo
//	@Tags		Todo
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Todo ID"
//	@Success	201	{object}	database.Response
//	@Router		/todo/{id}/complete [delete]
//	@Security	Bearer
func UnmarkComplete(ctx *gin.Context) {
	todo := ctx.MustGet("todo").(models.Todo)

	data := models.Todo{
		Status: enums.IN_PROGRESS,
	}

	response, err := database.Update(todo, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Todo updated"))

}
