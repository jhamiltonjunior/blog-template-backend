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
func (server *Server) routes() {
	// middleware := server.Router
	// middleware.Use()

	server.User()
	http.Handle("/register", server.Router)
	http.Handle("/user", server.Router)
}
