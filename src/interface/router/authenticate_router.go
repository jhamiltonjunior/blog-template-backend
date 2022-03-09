package router

import "github.com/jhamiltonjunior/priza-tech-backend/src/interface/controller"

func (server *Server) Authenticate() {
	auth := controller.Auth{}

	server.HandleFunc("/api/v1/authenticate", auth.Authenticate()).Methods("POST")
	
	// server.HandleFunc("/api/v1/authenticate/sso", listItem.CreateListItem()).Methods("POST")
}
