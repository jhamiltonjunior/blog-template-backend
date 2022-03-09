package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
	"github.com/jhamiltonjunior/priza-tech-backend/src/infra"
)

type List struct {
	ID     int    `json:"list_id" db:"list_id"`
	Title  string `json:"title" db:"title"`
	UserId int    `json:"user_id" db:"user_id"`
}

// This function will create the user by inserting his data into the database
//
// If everything goes correctly it will return a JSON
// with user data (minus password) and status code 201 created
// 
// Essa função vai criar o usuário inserindo os dados dele no banco de dados
// 
// Se tudo ocorrer corretamente ela irá retornar um JSON
// com os dados do usário (menos o password) e o status code 201 created
// 
func (list *List) CreateList() http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		json.NewDecoder(req.Body).Decode(list)

		_, err := infra.InsertList(
			`INSERT INTO list_schema (
				title, user_id
			)
			VALUES ($1, $2)
			RETURNING *`,
			list.Title, list.UserId,
		)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("%v", err),
			})

			return
		}

		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(list)
	}

}

// This function will show all lists referring to the id of the url
// /api/v{1}/list/{id:[0-9]+}
// 
// Esta função ira mostrar todas as listas referente ao id da url 
//  /api/v{1}/list/{id:[0-9]+}
// 
func (list *List) ShowList() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		sql := fmt.Sprintf("SELECT * FROM list_schema WHERE list_id=%v", params["id"])

		// row aqui está no singular pelo fata de que só existe um id para cada user
		// row here it is singular due to the fact that there is only one id for each user
		//
		row, err := config.Select(sql)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("Select List%v", err),
			})

			return
		}

		for row.Next() {
			err = row.Scan(
				&list.ID, &list.Title, &list.UserId,
			)
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(response).Encode(map[string]string{
					"message": fmt.Sprintf("Row scan: %v", err),
				})

				return
			}

			err = row.Close()
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(response).Encode(map[string]string{
					"message": fmt.Sprintf("Close row: %v", err),
				})

				return
			}
			json.NewEncoder(response).Encode(list)
		}
	}
}

// DeleteList will delete a list from the database
// and if it really works, it will return a simple JSON
// informed that the list has been deleted
// 
// DeleteList irá deletar uma lista do banco de dados
// e se realmente funcionar, irá retornar um simples JSON
// informado que a lista foi deletada
// 
func (list *List) DeleteList() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		sql := fmt.Sprintf("DELETE FROM list_schema WHERE list_id=%v", params["id"])

		_, err := config.Delete(sql)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"Fail": fmt.Sprintf("Error when deleting user: %v", err),
			})

			return
		}

		json.NewEncoder(response).Encode(map[string]string{
			"message": "List deleted with success!",
		})
	}
}
