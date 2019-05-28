package config

import (
	//"fmt"
)

// Configuration has the environment variables
type Configuration struct {
	HmacSecret string `yaml:"hmacsecret"`
	HmacBlockKey string `yaml:"hmacblockkey"`
	Linkedin OAuthConfig `yaml:"linkedin"`
}

// OAuthConfig has the credentials for authentication and authorization
type OAuthConfig struct {
	ID          string   `yaml:"clientid"`
	Secret      string   `yaml:"clientsecret"`
	AuthURL     string   `yaml:"auth_url"`
	ConnectURL  string   `yaml:"connect_url"`
	RedirectURL string   `yaml:"redirect_url"`
	TokenURL    string   `yaml:"token_url"`
	Scopes      []string `yaml:"scopes"`
}

// GetData parses yaml config
func (c *Configuration) GetData(env string) error {
	if err := parseFile(env, c); err != nil {
		return err
	}
	return nil
}

// GetConfig parses environment config into a struct
func GetConfig(env string, cf Data) error {
	if err := cf.GetData(env); err != nil {
		return err
	}
	return nil
}
