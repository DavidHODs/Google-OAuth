package controllers

import (
	"github.com/DavidHODs/Google-OAuth/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// init loads the neccessary configuration details required by oauth2 package.
func init() {
	utils.OAuthgolang = &oauth2.Config{
		RedirectURL:  utils.LoadFile("REDIRECT_URL"),
		ClientID:     utils.LoadFile("CLIENT_ID"),
		ClientSecret: utils.LoadFile("CLIENT_SECRET"),
		// scopes limits the access given to a token. this scope returns just the user info of the
		// signed in email address
		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint, //Endpoint is Google's OAuth 2.0 default endpoint
	}
}
