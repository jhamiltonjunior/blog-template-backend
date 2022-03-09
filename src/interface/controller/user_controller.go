package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
	"github.com/jhamiltonjunior/priza-tech-backend/src/infra"
)

// The User struct is responsible for getting the req.Body and inserting it into the database
// and the same is responsible for "porpulating" the JSON that returns from the database
//
// Please you from the frontend, redirect the user to the route
//  /api/v{n}/authenticate
// Here I just create the user, I don't have any JWT authenticate here
//
//
// User é o responsavel por pegar o req.Body e inserir no banco de dados
// e o mesmo é responsavel por "porpular" o JSON que retorna do banco de dados
//
// Por favor, você do frontend, redirecione o usuário para a rota
//  /api/v{n}/authenticate
// Aqui eu somente crio o usuário, não tenho nenhum JWT authenticate aqui
//
type User struct {
	ID int `json:"user_id" db:"user_id"`

	// I put Name, because if I put UserName when going to use
	// would have to call user.UserName and I don't like that
	// user.Name is already implied
	//
	// Coloquei Name, porque se eu colocasse UserName quando fosse usar
	// iria ter que chamar user.UserName e eu não gosto disso
	//  user.Name já fica subentendido
	Name      string    `json:"username" db:"username"`
	FullName  string    `json:"fullname" db:"fullname"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"passwd" db:"passwd"`
	CreatedAt string    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateUser There is an error here in returning user data
// it doesn't show the ID correctly, nor the Creation Date
// even though the user was created
// if you go in the route that shows all users you will see that
// he was raised
// Don't worry
// this could be changed in new feature
//
//  "user_id": 0,
//  "created_at": "",
//  "updated_at": ""
//
// Existe um erro aqui no retorno dos dados do usuario
// ele não mostra o ID corretamente, nem a Data de criação
// mesmo que o user foi criado
// se você for na rota que mostra todos os usuarios você vai ver que
// ele foi criado
// Não se preocupe
// Isso poderia ser mudado em uma nova feature
//
func (user *User) CreateUser() http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		json.NewDecoder(req.Body).Decode(user)

		_, err := infra.InsertUser(
			`INSERT INTO user_schema (
				username, fullname, email, passwd
			)
			VALUES ($1, $2, $3, $4)
			RETURNING *`,
			user.Name, user.FullName, user.Email, user.Password,
		)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("erro in Insert in create user: %v", err),
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
		user.Password = ""

		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(user)
	}

}

// ListAllUser will list ALL users that are registered in the database
// that's right, if there are 1 thousand users I advise you to put another 8 or 16 GB's of RAM
// on your machine
//
// In a new feature, a DESC LIMIT {NUMBER} OFFSET {NUMBER} could be placed in the query
// This will prevent server crashes or slowdowns.
//
// Can you imagine having to list 1000 users for 20 people at the same time?
//
// Consider being very careful with this.
//
// Why wasn't this implemented?
//
// This function shouldn't even exist!
// I ended up creating this function by accident, now it's a feature
//
//    <|Portugue Version|>
// ListAllUser vai listar TODOS os usários que estão cadastrados no banco de dados
// isso mesmo, se houver 1 mil usários eu aconselho você colocar mais uns 8 ou 16GB's de RAM
// na sua máquina
//
//  Em uma nova feature poderia ser colocado um DESC LIMIT {NUMBER} OFFSET {NUMBER} na query
// Isso ira impedir travamentos ou lentidão no servidor.
//
//  Imagina ter que listar 1 mil usuários para 20 pessoas ao mesmo tempo?
//
// Considere tomar bastante cuidado com isso.
// 
// Porque isso não foi implementado?
// 
//  Essa função nem deveria exitir!
// Acabei criando essa função por acidente, agora é uma feature
func (user *User) ListAllUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		rows, err := config.Select("SELECT * FROM user_schema")

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("erro in select query, in select all users: %v", err),
			})

			return
		}

		for rows.Next() {
			err = rows.Scan(
				&user.ID, &user.Name, &user.FullName, &user.Email,
				&user.Password, &user.CreatedAt, &user.UpdatedAt,
			)
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(response).Encode(map[string]string{
					"message": fmt.Sprintf("erro in row scan: %v", err),
				})

				return
			}

			// I'm putting the "", to overwrite password,
			// and don't display it to the end user
			// please do not use this in frontend application
			//
			// Eu estou colocando o "", para sobrescrever o password,
			// e não exibir para ao usuário final
			// por favor não use isso no frontend
			user.Password = ""

			if err := json.NewEncoder(response).Encode(user); err != nil {
				response.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(response).Encode(map[string]string{
					"message": fmt.Sprintf("erro in new encode json: %v", err),
				})

				return
			}
		}

		err = rows.Close()
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("erro in close rows: %v", err),
			})

			return
		}
	}
}

// ListUniqueUser Wil list a single user by id of url
//  /api/v{1}/user/{id:[0-9]+}
// If there is no error it will return a JSON with the referring user
// to the id of the url
// 
// ListUniqueUser Vai listar um unico usuário pelo id da url
//  /api/v{1}/user/{id:[0-9]+}
// 
// Se não houver nenhum erro irá retornar um JSON com o usuário referente
// ao id da url
func (user *User) ListUniqueUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		sql := fmt.Sprintf("SELECT * FROM user_schema WHERE user_id=%v", params["id"])

		// row aqui está no singular pelo fata de que só existe um id para cada user
		// row here it is singular due to the fact that there is only one id for each user
		//
		row, err := config.Select(sql)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("erro in select query, in select unique user: %v", err),
			})

			return
		}

		for row.Next() {
			err = row.Scan(
				&user.ID, &user.Name, &user.FullName, &user.Email,
				&user.Password, &user.CreatedAt, &user.UpdatedAt,
			)
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(response).Encode(map[string]string{
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

			err = json.NewEncoder(response).Encode(user)
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(response).Encode(map[string]string{
					"message": fmt.Sprintf("erro in new encode json: %v", err),
				})

				return
			}

			err = row.Close()
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(response).Encode(map[string]string{
					"message": fmt.Sprintf("erro in close row: %v", err),
				})

				return
			}
		}
	}
}

// This function will update the user data
// I was using insomnia and when I updated user data 1
// it was no longer listed at the beginning of the function
//
//  The last user to be modified goes to the end of ListAll()
//  At least that is how it was for me using *Insomnia*
//
// Essa função vai atualizar os dados do usuário
// eu estava usando o insomnia e quando eu atualizei o dado do user 1
// ele não era mais listado no inicio da função
//
//  O ultimo user a ser modificado vai para o final da ListAll()
//  Pelo menos foi assim comigo usando o Insomnia
//
func (user *User) UpdateUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		sql := fmt.Sprintf(`
			UPDATE user_schema
      SET
      username = $1,
      fullname = $2,
      email = $3,
      passwd = $4,
			updated_at = $5
      WHERE user_id = %v RETURNING *
		`, params["id"])

		json.NewDecoder(request.Body).Decode(user)

		user.UpdatedAt = time.Now()

		_, err := infra.UpdateUser(
			sql,
			user.Name, user.FullName, user.Email, user.Password, user.UpdatedAt,
		)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("Erro in update user: %v", err),
			})

			return
		}

		json.NewEncoder(response).Encode(user)
	}
}

// Will delete a user by id
// 
// Vai deletar um usuário pelo id
func (user *User) DeleteUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		sql := fmt.Sprintf("DELETE FROM user_schema WHERE user_id=%v", params["id"])

		_, err := config.Delete(sql)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"Fail": fmt.Sprintf("Error when deleting user: %v", err),
			})

			return
		}

		json.NewEncoder(response).Encode(map[string]string{
			"message": "User deleted with success!",
		})
	}
}
