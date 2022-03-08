package router

import (
	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/controller"
)

func (server *Server) List() {
	list := controller.List{}

	server.HandleFunc("/api/v1/list", list.CreateList()).Methods("POST")

	server.HandleFunc("/api/v1/list/{id:[0-9]+}", list.ShowList()).Methods("GET")
	server.HandleFunc("/api/v1/list/{id:[0-9]+}", list.DeleteList()).Methods("DELETE")
}