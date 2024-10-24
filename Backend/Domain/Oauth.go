package domain

import "golang.org/x/oauth2"

type OAuthConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUrl  string `json:"redirect_url"`
}

type OauthConfigInterface interface {
	InitialConfig() (*oauth2.Config, error)
}

type OauthUseCaseInterface interface {
	GoogleAuth() (interface{}, bool)
	GoogleCallback(code string) (User, *ErrorResponse, bool)
}