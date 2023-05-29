package routes

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/todo/configs"
	"github.com/rohanshrestha09/todo/controllers/sso"
)

// ShowAccount godoc
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	model.Account
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts/{id} [get]

func ssoRouter(router *gin.RouterGroup) {

	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))

	router.Use(sessions.Sessions("auth-session", store))

	router.GET("/login/facebook", sso.InitLogin(configs.GetFacebookOAuthConfig()))

	router.GET("/facebook/callback", sso.HandleLogin(configs.GetFacebookOAuthConfig()))

	router.GET("/login/google", sso.InitLogin(configs.GetGoogleOAuthConfig()))

	router.GET("/google/callback", sso.HandleLogin(configs.GetGoogleOAuthConfig()))

}
