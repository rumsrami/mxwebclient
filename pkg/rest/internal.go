package rest

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rumsrami/mxwebclient/pkg/config"
)

// getTLSServer returns an http server configured with TLS
func getTLSServer(env string, port string, r *mux.Router) (*http.Server, error) {
	certPath, err := config.GetPath("/assets/certs/")
	if err != nil {
		return nil, fmt.Errorf("cannot get server certificate, %v", err)
	}
	certificate, err := tls.LoadX509KeyPair(certPath+env+".server.crt", certPath+env+".server.key")
	if err != nil {
		return nil, fmt.Errorf("cannot get tls certificate, %v", err)
	}
	tlscfg := &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		InsecureSkipVerify: true,
	}

	tlsServer := &http.Server{
		TLSConfig:    tlscfg,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		Handler:      r,
		Addr:         ":" + port,
	}
	return tlsServer, nil
}

// func readConfig returns the configuration for the server
// returns TLS certificates aswell
func readConfig(env string) (*config.Configuration, error) {
	var cfg config.Configuration

	if err := config.GetConfig(env, &cfg); err != nil {
		return nil, fmt.Errorf("Error obtaining config, %v", err)
	}
	return &cfg, nil
}
