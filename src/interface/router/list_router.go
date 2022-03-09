package router

import (
	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/controller"
)

// List() is responsible for managing all List
//
// All routes that exist or that will exist must be here,
// It is also part of the Server struct
// 
// List() é responsavel por administrar todas Lista
// 
// Todas rotas que exitem ou que irão existir devem ficar aqui,
// Ela tambem faz parte do Server struct
// 
func (server *Server) List() {
	list := controller.List{}

	server.HandleFunc("/api/v1/list", list.CreateList()).Methods("POST")

	server.HandleFunc("/api/v1/list/{id:[0-9]+}", list.ShowList()).Methods("GET")
	server.HandleFunc("/api/v1/list/{id:[0-9]+}", list.DeleteList()).Methods("DELETE")
}