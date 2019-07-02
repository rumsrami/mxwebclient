// Routing for rami.im
// -> "portfolio.rami.im /" : static files or react for landing page
// -> "api.rami.im / <project name> /" : GRPC calls
// -> "dev.rami.im / <project name> /" : Clients for calling api

package rest

import (
	"github.com/gorilla/csrf"
	"github.com/rumsrami/mxwebclient/pkg/rand"
)

// Routes mounts the handlers on the router
func (s *HTTPServer) Routes() {
	b, _ := rand.Bytes(32)
	CSRF := csrf.Protect(b)
	s.r.Use(CSRF)
	s.r.PathPrefix("/index/").Handler(s.handleIndex())
	s.r.Handle("/mxweb/oauth2/linkedin/connect", s.HandleConnect())
	s.r.HandleFunc("/mxweb/oauth2/linkedin/callback", s.HandleCallback())
}
