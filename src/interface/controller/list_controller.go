package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/priza-tech-backend/src/infra"
)

type List struct {
	ID        int       `json:"list_id" db:"list_id"`
	Title     string    `json:"title" db:"title"`
	Checked   bool      `json:"checked" db:"checked"`
	UserId    int       `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (list *List) CreateList() http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		response.Header().Set("Content-Type", "application/json")

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

		json.NewEncoder(response).Encode(list)
	}

}

func (list *List) ShowList() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-type", "application/json")

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
				&list.ID, &list.Title, &list.Checked,
				&list.UserId, &list.CreatedAt, &list.UpdatedAt,
			)
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(response).Encode(map[string]string{
					"message": fmt.Sprintf("Row scan: %v", err),
				})
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
