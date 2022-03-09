package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
)

type Auth struct {
	UserName string `json:"username" db:"username"`
	Password string `json:"passwd" db:"passwd"`
}



// colocar as alterações feitas aqui no README

func (auth *Auth) Authenticate() http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		json.NewDecoder(req.Body).Decode(auth)

		sql := fmt.Sprintf(
			`SELECT * FROM user_schema
			WHERE username = '%v' AND passwd = '%v'`,
			auth.UserName, auth.Password,
		)

		_, err := config.Select(sql)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("%v", err),
			})

			// the return after the error the application continues that prevents
			//
			// o return impede que após o erro a aplização continue executando
			return
		}

		// I'm putting the "", to overwrite password,
		// and don't display it to the end user
		// please do not use this in frontend application
		//
		// Eu estou colocando o "", para sobrescrever o password,
		// e não exibir para ao usuário final
		// por favor não use isso no frontend
		auth.Password = ""

		json.NewEncoder(response).Encode(auth)
	}

}
