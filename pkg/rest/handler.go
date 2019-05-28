package rest

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/csrf"
	"github.com/rumsrami/mxwebclient/pkg/config"
)

// Auth handlers

// HandleConnect returns a HandlerFunc for logging in
// Calls Authorization URL
func (s *HTTPServer) HandleConnect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := csrf.Token(r)
		payload := map[string]string {
			"state" : state,
		}
		url := s.oauthClient.Oauth2.AuthCodeURL(state)

		secureCookie, err := s.Cookie.Encode("oauth_state", payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("Secure cookie unavailable: %v", err)
		}
		http.SetCookie(w, secureCookie)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

// HandleCallback returns a token
// Calls Callback URL and returns Token if state matches
func (s *HTTPServer) HandleCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("Unable to parse form: %v", err)
		}
		state := r.FormValue("state")
		//fmt.Printf("state: %v\n", state)
		payload := map[string]string {
			"state" : state,
		}
		//fmt.Printf("payload: %v\n", payload)
		cookie, err := r.Cookie("oauth_state")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Fatalf("Unable to get cookie value: %v", err)
		}
		//fmt.Printf("cookie: %v\n", cookie)
		valid, err := s.Cookie.Validate("oauth_state",cookie.Value, payload)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Fatalf("Unauthorized State: %v", err)
		}
		if valid {
			token, err := s.oauthClient.Oauth2.Exchange(r.Context(), r.FormValue("code"))
				if err != nil {
					fmt.Fprintf(w, "Unable to retrieve token: %v", err)
				}
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "token: %v", token)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

func (s *HTTPServer) handleIndex() http.Handler {
	execPath, err := config.GetExecPath()
	if err != nil {
		log.Fatalf("Cannot get path for static files: %v", err)
	}
	assetHandler := http.FileServer(http.Dir(filepath.Dir(execPath) + "/assets/static/"))
	// StripPrefix the subrouter + the path
	assetHandler = http.StripPrefix("/mxweb/index/", assetHandler)
	return assetHandler
}
