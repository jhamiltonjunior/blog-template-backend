package router

import (
	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/controller"
)

// This function will manage user routes
//
// Note that User is part of Server struct but there is no struct here
// this happens because User is part of package router
// and struct Server is part of package router
//
//  User() is package router and Server struct too
//
// Note que Usuário faz parte de Server struct, mas não tem nenhum struct aqui
// isso acontece pelo fato de que o User faz parte do packege router
// e a struct Serve também faz parte do package router
//
func (server *Server) User() {

	user := controller.User{}

	server.HandleFunc("/api/v1/user", user.ListAllUsers()).Methods("GET")
	server.HandleFunc("/api/v1/user", user.CreateUser()).Methods("POST")

	server.HandleFunc("/api/v1/user/{id:[0-9]+}", user.ListUser()).Methods("GET")
	server.HandleFunc("/api/v1/user/{id:[0-9]+}", user.UpdateUser()).Methods("PUT")

	// server.HandleFunc("/api/v1/user/{id}", user.CreateUser()).Methods("DELETE")
}
