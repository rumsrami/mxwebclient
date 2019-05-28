package main

import (
	"log"
	"os"

	"github.com/rumsrami/mxwebclient/pkg/rest"
)

var (
	port string
	env  string
)

func main() {

	port = os.Getenv("port")
	if port == "" {
		port = "443"
	}
	env = os.Getenv("env")
	if env == "" {
		env = "dev"
	}

	if srv, err := rest.NewHTTPServer(env, port); err != nil {
		log.Fatalf("Error, %v", err)
	} else {
		srv.Routes()
		srv.Server.ListenAndServeTLS("", "")
	}
}
