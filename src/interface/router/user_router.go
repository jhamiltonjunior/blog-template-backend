package router

import (
	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/controller"
)

// This function will manage user routes
//
func (server Server)User() {
	// var server = Server{
	// 	Router: mux.NewRouter(),
	// }

	user := controller.User{
		Name:     "Hamilton",
		Email:    "jose@gmail.com",
		Password: "1234",
	}

	// server.HandleFunc("/register", user.CreateUser()).Methods("POST")
	server.HandleFunc("/user", user.ListUser()).Methods("GET")
}
