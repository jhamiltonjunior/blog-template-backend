// A simple API  for TODO list
// This file is part of Priza Tech Back end API.
//
// You are free to modify and share this project or its files.
//
// Author   José Hamilton <https://github.com/jhamiltonjunior>
//
// Contact <https://www.linkedin.com/in/jhamiltonjunior>
// 
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/router"
	"github.com/joho/godotenv"
)

var (
	// godotenv vai fazer o arquivo .env funcionar
	// sem ele o os.Getenv não funciona
	err  = godotenv.Load(".env")

	// os.Getenv irá buscar o HTTP_PORT dentro do arquivo .env
	port = os.Getenv("HTTP_PORT")
)

func main() {
	// Essa captura de erro é referente ao godotenv
	if err != nil {
		fmt.Println(err)
	}

	router.NewServer()

	fmt.Printf("server listen in port%s \n", port)

	// Você irá encontrar o valor de port no seu arquivo .env ou .env.exemple
	http.ListenAndServe(port, nil)
}
