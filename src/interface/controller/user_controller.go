package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"passwd"`
}

func (user *User) CreateUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		config.Open(
			// "postgres://postgres@localhost/vibbra?sslmode=disable",
			"user=postgres password=0000 dbname=vibbra sslmode=disable",
		)

		config.Insert(
			"INSERT INTO user_schema (username, email, passwd) VALUES($1, $2, $3)",
			[]string{"Hamilton", "Jose Hamilton", "123"},
		)

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

func (user *User) ShowUser() http.HandlerFunc {
	// done := make(chan string)
	return func(writer http.ResponseWriter, req *http.Request) {
		result, err := config.Select("SELECT * FROM user_schema")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			writer.Header().Set("Content-type", "application/json")
			
			fmt.Println(err)
			
			return
		}
		fmt.Println(result)

		// column, err := result.Columns()
		// if err != nil {
		// 	fmt.Println(err)
			
		// 	return
		// }

		writer.Header().Set("Content-type", "application/json")
		if err := json.NewEncoder(writer).Encode(result); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)

			return
		}
	}
}
