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
	server.Routes()

	return server
}

// routes está em minusculo porque eu não preciso exportar
//
func (server *Server) Routes() {
	server.User()

	server.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		writer.Write([]byte("Hello"))
	}).Methods("GET")

	http.Handle("/", server.Router)
	http.Handle("/user", server.Router)
}
