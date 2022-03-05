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

	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/router"
	"github.com/joho/godotenv"
)

var (
	port = os.Getenv("HTTP_PORT")
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	if port == "" {
		port = ":1289"
	}

	router.NewServer()

	fmt.Printf("server listen in port%s", port)
	http.ListenAndServe(port, nil)
}
