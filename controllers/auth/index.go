package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rohanshrestha09/todo/configs"
	"github.com/rohanshrestha09/todo/enums"
	"github.com/rohanshrestha09/todo/service/database"

	"github.com/rohanshrestha09/todo/dto"
	"github.com/rohanshrestha09/todo/models"
	"github.com/rohanshrestha09/todo/utils"
)

// Auth godoc
//
//	@Summary	Get auth profile
//	@Tags		Auth
//	@Produce	json
//	@Success	200		{object}	database.GetResponse[models.User]
//	@Router		/auth 	[get]
//	@Security	Bearer
func Get(ctx *gin.Context) {
	authUser := ctx.MustGet("auth").(models.User)

	ctx.JSON(
		http.StatusOK,
		database.GetResponse[models.User]{
			Message: "Authorised",
			Data:    authUser,
		})

}

// Update Profile godoc
//
//	@Summary	Update profile
//	@Tags		Auth
//	@Accept		mpfd
//	@Produce	json
//	@Param		name	formData	string	false	"Name"
//	@Param		image	formData	file	false	"File to upload"
//	@Success	201		{object}	database.Response
//	@Router		/auth 	[patch]
//	@Security	Bearer
func Update(ctx *gin.Context) {

	authUser := ctx.MustGet("auth").(models.User)

	var profileUpdateDto dto.ProfileUpdateDTO

	if err := ctx.ShouldBindWith(&profileUpdateDto, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	var imageUrl, imageName string

	if file, err := ctx.FormFile("image"); err == nil {
		imageUrl, imageName, err = utils.UploadFile(file, enums.USER_DIR, enums.IMAGE)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if err := utils.DeleteFile(string(enums.USER_DIR) + authUser.ImageName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	data := models.User{
		Name:      profileUpdateDto.Name,
		Image:     imageUrl,
		ImageName: imageName,
	}

	response, err := database.Update(authUser, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Profile Updated"))

}

// Delete Profile godoc
//
//	@Summary	Delete profile
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Success	201		{object}	database.Response
//	@Router		/auth 	[delete]
//	@Security	Bearer
func Delete(ctx *gin.Context) {
	authUser := ctx.MustGet("auth").(models.User)

	if authUser.Image != "" && authUser.ImageName != "" {
		if err := utils.DeleteFile(string(enums.USER_DIR) + authUser.ImageName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	var todoFileNames []string

	tx := configs.DB.Begin()

	tx.Where(&models.Todo{UserID: authUser.ID}).Pluck("file_name", &todoFileNames)

	err := tx.Unscoped().Where(&models.Todo{UserID: authUser.ID}).Delete(&models.Todo{}).Error

	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = tx.Unscoped().Where(&models.List{UserID: authUser.ID}).Delete(models.List{}).Error

	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = tx.Unscoped().Delete(&authUser).Error

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

	ctx.JSON(http.StatusCreated, database.Response{Message: "User deleted"})
}
