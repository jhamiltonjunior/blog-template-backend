package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	server := &Server{
		Router: mux.NewRouter(),
	}

	server.routes()

	return server
}

// routes it's lowercase because I don't need to export
//
//rotas é minúscula porque não preciso exportar
//
func (server *Server) routes() {
	// middleware := server.Router
	// middleware.Use()

	// This is a gorilla/mux requirement
	// I need to pass the server.Router as the second parameter
	//
	// Isso aqui é um requisito do gorilla/mux
	// Eu preciso passar o server.Router como segundo parametro
	//
	server.User()
	http.Handle("/", server.Router)
}
