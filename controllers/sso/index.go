package sso

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/todo/configs"
	"github.com/rohanshrestha09/todo/enums"
	"github.com/rohanshrestha09/todo/models"
	"github.com/rohanshrestha09/todo/utils"
	"gorm.io/gorm"
)

func InitLogin(OAuth2Config *configs.OAuth2Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		state, err := configs.GetRandomOAuthStateString()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		session := sessions.Default(ctx)

		session.Set("state", state)

		if err := session.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		url := OAuth2Config.AuthCodeURL(state)

		ctx.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func HandleLogin(OAuth2Config *configs.OAuth2Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		db := configs.DB

		session := sessions.Default(ctx)

		if ctx.Query("state") != session.Get("state") {
			ctx.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		token, err := OAuth2Config.Exchange(context.TODO(), ctx.Query("code"))

		if err != nil || token == nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		userDetails, err := (func() (*models.SSOUser, error) {

			switch OAuth2Config.Provider {
			case enums.Facebook:
				userDetails, err := utils.GetSSOUserInfo[models.FacebookUser](token.AccessToken, OAuth2Config.TokenURI)
				return &models.SSOUser{
					ID:    userDetails.ID,
					Name:  userDetails.Name,
					Email: userDetails.Email,
					Image: userDetails.Picture.Data.Url,
				}, err

			case enums.Google:
				userDetails, err := utils.GetSSOUserInfo[models.GoogleUser](token.AccessToken, OAuth2Config.TokenURI)
				return &models.SSOUser{
					ID:    userDetails.ID,
					Name:  userDetails.Name,
					Email: userDetails.Email,
					Image: userDetails.Picture,
				}, err

			default:
				return &models.SSOUser{}, errors.New("invalid provider")
			}

		})()

		if err != nil {
			ctx.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
		}

		var user models.User

		user.Email = userDetails.Email

		err = db.Where(&models.User{Email: user.Email}).First(&models.User{}).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {

				user.Name = userDetails.Name

				user.Image = userDetails.Image

				user.Password = userDetails.ID

				if err := db.Create(&user).Error; err != nil {
					ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
					return
				}
			}

			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
			return
		}

		authToken, err := utils.SignJwt(user.Email)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.SetCookie("token", authToken, 30*1400*60, "/", "localhost", false, true)

		ctx.Redirect(http.StatusTemporaryRedirect, "/profile")
	}
}
