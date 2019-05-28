package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rumsrami/mxwebclient/pkg/auth"
	"github.com/rumsrami/mxwebclient/pkg/config"
	"github.com/rumsrami/mxwebclient/pkg/secure"
)

// HTTPServer represents all components of the service
type HTTPServer struct {
	Config      config.Configuration
	Cookie      secure.Cookie
	Server      *http.Server
	oauthClient auth.Client
	r           mux.Router
}

// NewHTTPServer returns an HTTPServer for routes mouting and launching
// Called in the main package
func NewHTTPServer(env string, port string) (*HTTPServer, error) {
	var err error
	s := HTTPServer{}

	// HTTPServer Config
	cfg, err := readConfig(env)
	if err != nil {
		return nil, err
	}
	s.Config = *cfg

	// HTTPServer Router
	r := mux.NewRouter().StrictSlash(true)
	// Subrouter should not have a trailing slash
	mxclient := r.PathPrefix("/mxweb").Subrouter()
	s.r = *mxclient

	// HTTPServer TLS Server
	s.Server, err = getTLSServer(env, port, &s.r)
	if err != nil {
		return nil, err
	}

	// HTTPServer Auth Client
	authClientConfig := auth.ClientConfig{
		ID:          cfg.Linkedin.ID,
		Secret:      cfg.Linkedin.Secret,
		AuthURL:     cfg.Linkedin.AuthURL,
		ConnectURL:  cfg.Linkedin.ConnectURL,
		RedirectURL: cfg.Linkedin.RedirectURL,
		TokenURL:    cfg.Linkedin.TokenURL,
		Scopes:      cfg.Linkedin.Scopes,
	}
	linkedinAuthClient := auth.NewAuthClient(authClientConfig)
	s.oauthClient = auth.Client{
		Oauth2: linkedinAuthClient.Oauth2,
	}

	// HTTPServer Secure Cookie
	cookieConfig := secure.CookieConfig{
		Hmac:      cfg.HmacSecret,
		HmacBlock: cfg.HmacBlockKey,
	}
	cookie, err := secure.NewCookie(cookieConfig)
	if err != nil {
		log.Fatalf("Secure cookie unavailable: %v", err)
	}
	s.Cookie = *cookie

	return &s, nil
}
