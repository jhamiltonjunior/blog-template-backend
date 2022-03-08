package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/priza-tech-backend/src/infra"
)

type List struct {
	ID     int    `json:"list_id" db:"list_id"`
	Title  string `json:"title" db:"title"`
	UserId int    `json:"user_id" db:"user_id"`
}

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

func (list *List) ShowList() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		sql := fmt.Sprintf("SELECT * FROM list_schema WHERE list_id=%v", params["id"])

		// row aqui está no singular pelo fata de que só existe um id para cada user
		// row here it is singular due to the fact that there is only one id for each user
		//
		row, err := infra.SelectList(sql)
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

func (list *List) DeleteList() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		sql := fmt.Sprintf("DELETE FROM list_schema WHERE list_id=%v", params["id"])

		_, err := infra.DeleteList(sql)
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
