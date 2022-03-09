package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/middleware"
)

type Server struct {
	*mux.Router
}

// This function is responsible for starting the server
// it will contain the instance of the struct Server
// and calling the routes server.routes()
// 
// Esta função é responsavel por iniciar o servidor
// nela irá conter a instância da struct Server
//  e a chamada das rotas server.routes()
// 
func NewServer() *Server {
	server := &Server{
		Router: mux.NewRouter(),
	}

	server.routes()

	return server
}

// routes is responsible for calling all other routes
// it is also responsible for calling the middleware that it will do
// make all other routes have a Content-type: application/json
// set in response header
//
// That is, it will always return a JSON
func (server *Server) routes() {
	middlewares := server.Router
	middlewares.Use(middleware.SetContentType)
	
	server.User()
	server.Authenticate()
	server.List()
	server.ListItem()

	// This is a gorilla/mux requirement
	// I need to pass the server.Router as the second parameter
	//
	http.Handle("/", server.Router)
}
