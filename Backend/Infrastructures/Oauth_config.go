package infrastructures

import (
	"errors"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Oauth struct{
	config *Config
}

func NewOauthConfig(config *Config) *Oauth{
	return &Oauth{
		config: config,
	}
}

func (o *Oauth) InitialConfig() (*oauth2.Config, error) {
	if o.config.ClientId == "" || o.config.ClientSecret == "" || o.config.RedirectUrl == "" {
		return nil, errors.New("missing oauth initial configuration")
	}
	return &oauth2.Config{
		ClientID:     o.config.ClientId,
		ClientSecret: o.config.ClientSecret,
		RedirectURL:  o.config.RedirectUrl,
		Scopes : []string{"email", "profile", "openid"},
		Endpoint: google.Endpoint,
	}, nil
}
