// TODO fix errors to http error instead of log and fmt

package auth

import (
	"golang.org/x/oauth2"
)

// Client represents 3rd party Authentication providers
type Client struct {
	Oauth2       oauth2.Config
}

// ClientConfig Configures an OauthClient
type ClientConfig struct {
	ID          string   
	Secret      string   
	AuthURL     string   
	ConnectURL  string   
	RedirectURL string   
	TokenURL    string   
	Scopes      []string
}

// NewAuthClient returns a new 3rd party Authenticator
func NewAuthClient(config ClientConfig) (*Client) {
	oauth2Client := oauth2.Config{
		ClientID:     config.ID,
		ClientSecret: config.Secret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.AuthURL,
			TokenURL: config.TokenURL,
		},
		RedirectURL: config.RedirectURL,
		Scopes:      config.Scopes,
	}
	return &Client {
		Oauth2: oauth2Client,
	}
}
