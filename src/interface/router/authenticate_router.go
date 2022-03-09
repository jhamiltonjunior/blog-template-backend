package router

import "github.com/jhamiltonjunior/priza-tech-backend/src/interface/controller"

// Authenticate() is responsible for handling user authentication
//
// All routes that exist or that will exist must be here,
// It is also part of the Server struct
// 
// Authenticate() é responsavel por administrar a autenticação do usuário
//
// Todas rotas que exitem ou que irão existir devem ficar aqui,
// Ela tambem faz parte do Server struct
// 
func (server *Server) Authenticate() {
	auth := controller.Auth{}

	server.HandleFunc("/api/v1/authenticate", auth.Authenticate()).Methods("POST")

	server.HandleFunc("/api/v1/authenticate/sso", auth.AuthenticateSSO()).Methods("POST")
}
