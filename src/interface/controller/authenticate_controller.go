package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
)

type Auth struct {
	UserName string `json:"username" db:"username"`
	Password string `json:"passwd" db:"passwd"`
}

// Uma função fake que retorna o dado inserido
// no req.Body e um token que também é fake
//
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

		token, err := getJWT()
		if err != nil {
			json.NewEncoder(response).Encode(err)

			return
		}

		json.NewEncoder(response).Encode(auth)
		json.NewEncoder(response).Encode(map[string]string{
			"token": fmt.Sprint(token),
		})
	}

}

// Assim como a função anterior essa aqui não é diferente,
// ela vai retorna o dado inserido
// no req.Body e um token que também é fake
func (auth *Auth) AuthenticateSSO() http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		json.NewDecoder(req.Body).Decode(auth)

		sql := fmt.Sprintf(
			`SELECT * FROM user_schema
			WHERE username = '%v'`,
			auth.UserName,
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

		token, err := getJWT()
		if err != nil {
			json.NewEncoder(response).Encode(err)

			return
		}

		json.NewEncoder(response).Encode(auth)
		json.NewEncoder(response).Encode(map[string]string{
			"token": fmt.Sprint(token),
		})
	}

}

func getJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Krissanawat"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ...
