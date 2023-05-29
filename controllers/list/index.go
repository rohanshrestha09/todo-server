package list

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/todo/configs"
	"github.com/rohanshrestha09/todo/dto"
	"github.com/rohanshrestha09/todo/enums"
	"github.com/rohanshrestha09/todo/models"
	"github.com/rohanshrestha09/todo/service/database"
	"github.com/rohanshrestha09/todo/utils"
)

// Get List godoc
//
//	@Summary	Get a list
//	@Tags		List
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"List ID"
//	@Success	200	{object}	database.GetResponse[models.List]
//	@Router		/list/{id} [get]
//	@Security	Bearer
func Get(ctx *gin.Context) {

	list := ctx.MustGet("list").(models.List)

	ctx.JSON(
		http.StatusOK,
		database.GetResponse[models.List]{
			Message: "List Fetched",
			Data:    list,
		})

}

// Get List godoc
//
//	@Summary	Get all lists
//	@Tags		List
//	@Accept		json
//	@Produce	json
//	@Param		page	query		int		false	"Page"
//	@Param		size	query		int		false	"Page size"
//	@Param		sort	query		string	false	"Sort"	Enums(id, created_at, name)
//	@Param		order	query		string	false	"Order"	Enums(asc, desc)
//	@Param		search	query		string	false	"Search"
//	@Success	200		{object}	database.GetAllResponse[models.List]
//	@Router		/list [get]
//	@Security	Bearer
func GetAll(ctx *gin.Context) {

	authUser := ctx.MustGet("auth").(models.User)

	args := database.GetAllArgs[models.List]{
		Pagination: true,
		Search:     true,
		Filter: models.List{
			UserID: authUser.ID,
		},
	}

	response, err := database.GetAll(ctx.BindQuery, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// Create List godoc
//
//	@Summary	Create a list
//	@Tags		List
//	@Accept		mpfd
//	@Produce	json
//	@Param		body	body		dto.ListCreateDTO	true	"Request Body"
//	@Success	201		{object}	database.Response
//	@Router		/list [post]
//	@Security	Bearer
func Create(ctx *gin.Context) {
	authUser := ctx.MustGet("auth").(models.User)

	var listCreateDto dto.ListCreateDTO

	if err := ctx.BindJSON(&listCreateDto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	project := models.List{
		UserID: authUser.ID,
		Name:   listCreateDto.Name,
	}

	response, err := database.Create(&project)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("List Created"))
}

// Update List godoc
//
//	@Summary	Update a list
//	@Tags		List
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int					true	"List ID"
//	@Param		body	body		dto.ListUpdateDTO	true	"Request Body"
//	@Success	201		{object}	database.Response
//	@Router		/list/{id} [patch]
//	@Security	Bearer
func Update(ctx *gin.Context) {
	list := ctx.MustGet("list").(models.List)

	var listUpdateDto dto.ListUpdateDTO

	if err := ctx.BindJSON(&listUpdateDto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	data := models.List{
		Name: listUpdateDto.Name,
	}

	response, err := database.Update(&list, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("List Updated"))
}

// Delete List godoc
//
//	@Summary	Delete a list
//	@Tags		List
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"List ID"
//	@Success	201	{object}	database.Response
//	@Router		/list/{id} [delete]
//	@Security	Bearer
func Delete(ctx *gin.Context) {
	list := ctx.MustGet("list").(models.List)

	var todoFileNames []string

	tx := configs.DB.Begin()

	tx.Where(&models.Todo{ListID: list.ID}).Pluck("file_name", &todoFileNames)

	err := tx.Unscoped().Where(&models.Todo{ListID: list.ID}).Delete(&models.Todo{}).Error

	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = tx.Unscoped().Delete(&list).Error

	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	tx.Commit()

	go func(todoFileNames []string) {
		for _, filename := range todoFileNames {
			utils.DeleteFile(string(enums.TODO_DIR) + filename)
		}
	}(todoFileNames)

	ctx.JSON(http.StatusCreated, database.Response{Message: "List Deleted"})
}
