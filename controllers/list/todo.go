package list

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/todo/models"
	"github.com/rohanshrestha09/todo/service/database"
)

// Get Todo godoc
//
//	@Summary	Get todos
//	@Tags		List
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int		true	"List ID"
//	@Param		page	query		int		false	"Page"
//	@Param		size	query		int		false	"Page size"
//	@Param		sort	query		string	false	"Sort"	Enums(id, created_at, name)
//	@Param		order	query		string	false	"Order"	Enums(asc, desc)
//	@Param		search	query		string	false	"Search"
//	@Success	200		{object}	database.GetAllResponse[models.Todo]
//	@Router		/list/{id}/todo [get]
//	@Security	Bearer
func GetTodos(ctx *gin.Context) {
	list := ctx.MustGet("list").(models.List)

	args := database.GetAllArgs[models.Todo]{
		Pagination: true,
		Search:     true,
		Filter: models.Todo{
			ListID: list.ID,
		},
	}

	response, err := database.GetAll(ctx.BindQuery, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
