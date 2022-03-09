package router

import (
	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/controller"
)

// ListItem() is responsible for managing all List Items
//
// All routes that exist or that will exist must be here,
// It is also part of the Server struct
func (server *Server) ListItem() {
	listItem := controller.ListItem{}

	server.HandleFunc("/api/v1/list/{id:[0-9]+}/item", listItem.ShowListItem()).Methods("GET")
	server.HandleFunc("/api/v1/list/{id:[0-9]+}/item", listItem.CreateListItem()).Methods("POST")

	server.HandleFunc(
		"/api/v1/list/{id:[0-9]+}/item/{item_id:[0-9]+}", listItem.UpdateListItem(),
	).Methods("PUT")

	server.HandleFunc(
		"/api/v1/list/{id:[0-9]+}/item/{item_id:[0-9]+}", listItem.DeleteListItem(),
	).Methods("DELETE")
}
