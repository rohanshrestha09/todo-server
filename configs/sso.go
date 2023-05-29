package configs

import (
	"crypto/rand"
	"encoding/base64"
	"os"

	"github.com/rohanshrestha09/todo/enums"
	"golang.org/x/oauth2/google"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

const (
	FacebookURI = "https://graph.facebook.com/me?fields=id,name,email,picture&access_token="
	GoogleURI   = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

type OAuth2Config struct {
	*oauth2.Config
	Provider enums.Provider
	TokenURI string
}

func GetFacebookOAuthConfig() *OAuth2Config {
	return &OAuth2Config{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("FB_APP_ID"),
			ClientSecret: os.Getenv("FB_SECRET"),
			RedirectURL:  os.Getenv("FB_CALLBACK_URL"),
			Endpoint:     facebook.Endpoint,
			Scopes:       []string{"email"},
		},
		Provider: enums.Facebook,
		TokenURI: FacebookURI,
	}
}

func GetGoogleOAuthConfig() *OAuth2Config {
	return &OAuth2Config{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_APP_ID"),
			ClientSecret: os.Getenv("GOOGLE_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_CALLBACK_URL"),
			Endpoint:     google.Endpoint,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		},
		Provider: enums.Google,
		TokenURI: GoogleURI,
	}
}

// GetRandomOAuthStateString will return random string
func GetRandomOAuthStateString() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, err
}
