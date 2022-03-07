package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
)

// _, err = db.Exec(sql, values)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// user: 'postgres',
//   host: 'localhost',
//   database: 'nauts',
//   password: '0000',
//   port: 5432,
// "postgres://postgres@localhost/testdb?sslmode=disable"

type User struct {
	ID        int    `json:"user_id" db:"user_id"`
	Name      string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"passwd"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

func (user *User) CreateUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		body, _ := ioutil.ReadAll(req.Body)

		// fmt.Println(string(body))

		row, err := config.Insert(
			`INSERT INTO user_schema (
				username, fullname, email, passwd
			)
			VALUES ($1, $2, $3, $4)
			RETURNING *`,
			// []string{"hamilton", "José Hamilton", "jose@hamilton.com", "123"},
			"hamilton", "José Hamilton", "jose@hamilton.com", "123",
		)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(writer).Encode(map[string]string{
				"message": fmt.Sprintf("erro in Insert: %v", err),
			})

			return
		}

		// 	namedSQL := `INSERT INTO "clients" ("name","address") VALUES (:name,:address)`

		// _, err = db.NamedExec(namedSQL,
		// 	client{
		// 		Name:    "Cyberdyne Systems",
		// 		Address: "2144 Kramer St",
		// 	})

		json.NewEncoder(writer).Encode(row)
	}

}

func (user *User) ListAllUsers() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		rows, err := config.Select("SELECT * FROM user_schema")

		writer.Header().Set("Content-type", "application/json")

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(writer).Encode(map[string]string{
				"message": fmt.Sprintf("erro in Select query: %v", err),
			})

			return
		}

		for rows.Next() {
			user := User{}

			err = rows.Scan(
				&user.ID, &user.Name, &user.Email,
				&user.Password, &user.CreatedAt, &user.UpdatedAt,
			)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(writer).Encode(map[string]string{
					"message": fmt.Sprintf("erro in row scan: %v", err),
				})

				return
			}

			// I'm putting the "", to overwrite password,
			// and don't display it to the end user
			// please do not use this in frontend application
			//
			// Eu estou colocando o "", para sobrescrever o password,
			// e não exibir par ao usuário final
			// por favor não use isso no frontend
			user.Password = ""

			if err := json.NewEncoder(writer).Encode(user); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(writer).Encode(map[string]string{
					"message": fmt.Sprintf("erro in new encode json: %v", err),
				})

				return
			}
		}

		err = rows.Close()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(writer).Encode(map[string]string{
				"message": fmt.Sprintf("erro in close rows: %v", err),
			})

			return
		}
	}
}

func (user *User) ListUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		query := fmt.Sprintf("SELECT * FROM user_schema WHERE user_id=%v", params["user_id"])

		writer.Header().Set("Content-type", "application/json")

		// row aqui está no singular pelo fata de que só existe um id para cada user
		// row here it is singular due to the fact that there is only one id for each user
		//
		row, err := config.Select(query)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(writer).Encode(map[string]string{
				"message": fmt.Sprintf("erro in Select query: %v", err),
			})

			return
		}

		for row.Next() {
			err = row.Scan(
				&user.ID, &user.Name, &user.Email,
				&user.Password, &user.CreatedAt, &user.UpdatedAt,
			)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(writer).Encode(map[string]string{
					"message": fmt.Sprintf("erro in row scan: %v", err),
				})
			}

			// I'm putting the "", to overwrite password,
			// and don't display it to the end user
			// please do not use this in frontend application
			//
			// Eu estou colocando o "", para sobrescrever o password,
			// e não exibir par ao usuário final
			// por favor não use isso no frontend
			user.Password = ""

			err = json.NewEncoder(writer).Encode(user)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(writer).Encode(map[string]string{
					"message": fmt.Sprintf("erro in new encode json: %v", err),
				})

				return
			}

			err = row.Close()
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(writer).Encode(map[string]string{
					"message": fmt.Sprintf("erro in close row: %v", err),
				})

				return
			}
		}
	}
}

func (user *User) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
