package router

import (
	"github.com/google/uuid"
	"github.com/jhamiltonjunior/priza-tech-backend/src/interface/controller"
)

// This function will manage user routes
//
func (server Server) User() {
	myuuid, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	user := controller.User{
		ID:       myuuid,
		Name:     "Hamilton",
		Email:    "jose@gmail.com",
		Password: "1234",
	}

	server.HandleFunc("/register", user.CreateUser()).Methods("POST")
	server.HandleFunc("/user", user.ShowUser()).Methods("GET")
}
