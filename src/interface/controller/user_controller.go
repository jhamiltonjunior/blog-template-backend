package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

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

		// var err error
		// user.ID, err = uuid.NewUUID()
		// if err != nil {
		// 	panic(err)
		// }

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


		for result.Next() {
			user := User{}

			err = result.Scan(
				&user.ID, &user.Name, &user.Email,
				&user.Password, &user.CreatedAt, &user.UpdatedAt,
			)
			if err != nil {
				err = fmt.Errorf("erro in result.Scan %v", err)

				http.Error(writer, err.Error(), http.StatusInternalServerError)
				writer.Header().Set("Content-type", "application/json")

				return
			}

			user.Password = ""

			writer.Header().Set("Content-type", "application/json")
			if err := json.NewEncoder(writer).Encode(user); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)

				return
			}

			fmt.Println(user.ID)
			fmt.Println(user.Name)
			fmt.Println(user.Email)
			fmt.Println(user.CreatedAt)
			fmt.Println(user.UpdatedAt)
		}

		err = result.Close()
		if err != nil {
			err = fmt.Errorf("erro in result.Close %v", err)

			http.Error(writer, err.Error(), http.StatusInternalServerError)
			writer.Header().Set("Content-type", "application/json")

			return
		}

		writer.Header().Set("Content-type", "application/json")
		if err := json.NewEncoder(writer).Encode("result"); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)

			return
		}
	}
}
