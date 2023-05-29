package todo

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rohanshrestha09/todo/dto"
	"github.com/rohanshrestha09/todo/enums"
	"github.com/rohanshrestha09/todo/models"
	"github.com/rohanshrestha09/todo/service/database"
	"github.com/rohanshrestha09/todo/utils"
)

// Get Todo godoc
//
//	@Summary	Get a todo
//	@Tags		Todo
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Todo ID"
//	@Success	200	{object}	database.GetResponse[models.Todo]
//	@Router		/todo/{id} [get]
//	@Security	Bearer
func Get(ctx *gin.Context) {

	todo := ctx.MustGet("todo").(models.Todo)

	ctx.JSON(
		http.StatusOK,
		database.GetResponse[models.Todo]{
			Message: "Todo Fetched",
			Data:    todo,
		})

}

// Get Todo godoc
//
//	@Summary	Get all todos
//	@Tags		Todo
//	@Accept		json
//	@Produce	json
//	@Param		page	query		int		false	"Page"
//	@Param		size	query		int		false	"Page size"
//	@Param		sort	query		string	false	"Sort"	Enums(id, created_at, name)
//	@Param		order	query		string	false	"Order"	Enums(asc, desc)
//	@Param		search	query		string	false	"Search"
//	@Success	200		{object}	database.GetAllResponse[models.Todo]
//	@Router		/todo [get]
//	@Security	Bearer
func GetAll(ctx *gin.Context) {

	authUser := ctx.MustGet("auth").(models.User)

	args := database.GetAllArgs[models.Todo]{
		Pagination: true,
		Search:     true,
		Filter: models.Todo{
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

// Create Todo godoc
//
//	@Summary	Create a todo
//	@Tags		Todo
//	@Accept		mpfd
//	@Produce	json
//	@Param		id			path		int		true	"List ID"
//	@Param		name		formData	string	true	"Name"
//	@Param		due			formData	string	true	"Due Time"	format(dateTime)
//	@Param		important	formData	boolean	false	"Important"
//	@Param		note		formData	string	false	"Note"
//	@Param		file		formData	file	false	"File to upload"
//	@Success	201			{object}	database.Response
//	@Router		/todo/{id} [post]
//	@Security	Bearer
func Create(ctx *gin.Context) {
	authUser := ctx.MustGet("auth").(models.User)

	list := ctx.MustGet("list").(models.List)

	var todoCreateDto dto.TodoCreateDTO

	if err := ctx.ShouldBindWith(&todoCreateDto, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	todo := models.Todo{
		UserID:    authUser.ID,
		ListID:    list.ID,
		Name:      todoCreateDto.Name,
		Note:      todoCreateDto.Note,
		Start:     time.Now(),
		Due:       todoCreateDto.Due,
		Important: todoCreateDto.Important,
	}

	if file, err := ctx.FormFile("file"); err == nil {
		todo.File, todo.FileName, err = utils.UploadFile(file, enums.TODO_DIR, enums.IMAGE)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	response, err := database.Create(&todo)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Todo Created"))
}

// Update Todo godoc
//
//	@Summary	Update a Todo
//	@Tags		Todo
//	@Accept		mpfd
//	@Produce	json
//	@Param		id		path		int		true	"Todo ID"
//	@Param		name	formData	string	false	"Name"
//	@Param		due		formData	string	false	"Due Time"	format(dateTime)
//	@Param		note	formData	string	false	"Note"
//	@Param		file	formData	file	false	"File to upload"
//	@Success	201		{object}	database.Response
//	@Router		/todo/{id} [patch]
//	@Security	Bearer
func Update(ctx *gin.Context) {
	todo := ctx.MustGet("todo").(models.Todo)

	var todoUpdateDto dto.TodoUpdateDTO

	if err := ctx.ShouldBindWith(&todoUpdateDto, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	var fileUrl, fileName string

	if file, err := ctx.FormFile("image"); err == nil {
		fileUrl, fileName, err = utils.UploadFile(file, enums.TODO_DIR, enums.IMAGE)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if err := utils.DeleteFile(string(enums.TODO_DIR) + todo.FileName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	data := models.Todo{
		Name:     todoUpdateDto.Name,
		Note:     todoUpdateDto.Note,
		Due:      todoUpdateDto.Due,
		File:     fileUrl,
		FileName: fileName,
	}

	response, err := database.Update(&todo, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Todo Updated"))
}

// Delete Todo godoc
//
//	@Summary	Delete a todo
//	@Tags		Todo
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Todo ID"
//	@Success	201	{object}	database.Response
//	@Router		/todo/{id} [delete]
//	@Security	Bearer
func Delete(ctx *gin.Context) {
	todo := ctx.MustGet("todo").(models.Todo)

	if todo.File != "" && todo.FileName != "" {
		if err := utils.DeleteFile(string(enums.TODO_DIR) + todo.FileName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	response, err := database.Delete(&todo)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Todo Updated"))
}
