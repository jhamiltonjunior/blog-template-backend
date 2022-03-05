package controller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (user *User) CreateUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)

			return
		}

		var err error
		user.ID, err = uuid.NewUUID()
		if err != nil {
			panic(err)
		}

		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(user); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)

			return
		}
	}

}

func (user *User) ListUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-type", "application/json")

		if err := json.NewEncoder(writer).Encode(user); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)

			return
		}
	}
}
