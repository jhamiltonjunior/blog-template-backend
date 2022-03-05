// A simple API  for template of a blog
// This file is part of Priza Tech Back end API.
//
// You are free to modify and share this project or its files.
//
// Author   José Hamilton <https://github.com/jhamiltonjunior>
//
// Contact <https://www.linkedin.com/in/jhamiltonjunior>
//
// This code goes create outher file with the content to Web Scrapping
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/router"
)

var (
	port = os.Getenv("HTTP_PORT")
	// muxR = mux.NewRouter()
)

func main() {
	if port == "" {
		port = ":1289"
	}

	server := &router.Server{
		Router: mux.NewRouter(),
	}
	server.Routes()

	// router.NewServer()

	fmt.Printf("server listen in port%s", port)

	http.Handle("/", server.Router)
	http.Handle("/user", server.Router)
	// http.Handle("/register", server.Router)

	http.ListenAndServe(port, nil)
}
