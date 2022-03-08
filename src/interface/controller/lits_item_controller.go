package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/priza-tech-backend/src/infra"
)

type ListItem struct {
	ID          int        `json:"list_item_id" db:"list_item_id"`
	UserId      int        `json:"user_id" db:"user_id"`
	ListId      int        `json:"list_id" db:"list_id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	CreatedAt   string     `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

func (listItem *ListItem) CreateListItem() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		json.NewDecoder(request.Body).Decode(listItem)

		_, err := infra.InsertListItem(
			`INSERT INTO list_item_schema (
				user_id, list_id, title, description
			)
			VALUES ($1, $2, $3, $4)
			RETURNING *`,
			listItem.UserId, listItem.ListId, listItem.Title, listItem.Description,
		)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"error": fmt.Sprintf("%v", err),
			})

			return
		}

		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(listItem)
	}
}

func (listItem *ListItem) ShowListItem() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		items := []ListItem{}

		sql := fmt.Sprintf("SELECT * FROM list_item_schema WHERE list_id=%v", params["id"])

		// row aqui está no singular pelo fata de que só existe um id para cada user
		// row here it is singular due to the fact that there is only one id for each user
		//
		row, err := infra.SelectListItem(sql)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("Select List%v", err),
			})

			return
		}

		for row.Next() {
			err = row.StructScan(
				listItem,
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

			items = append(items, *listItem)

			json.NewEncoder(response).Encode(items)
		}
	}
}

func (listItem *ListItem) UpdateListItem() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

	}
}

func (listItem *ListItem) DeleteListItem() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

	}
}
