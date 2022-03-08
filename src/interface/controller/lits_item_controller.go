package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
	"github.com/jhamiltonjunior/priza-tech-backend/src/infra"
)

type ListItem struct {
	ID          int    `json:"list_item_id" db:"list_item_id"`
	UserId      int    `json:"user_id" db:"user_id"`
	ListId      int    `json:"list_id" db:"list_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
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
		db, err := infra.SelectListItem(sql)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("Select List%v", err),
			})

			return
		}

		err = db.Select(&items, sql)
		if err != nil {
			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf("%v", err),
			})

			return
		}

		json.NewEncoder(response).Encode(items)
	}
}

func (listItem *ListItem) UpdateListItem() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		fmt.Println(params["item_id"])

		sql := fmt.Sprintf(`
			UPDATE list_item_schema
      SET
      title = $1,
      description = $2,
      user_id = $3,
      list_id = $4
      WHERE list_id = %v AND list_item_id = %v
			RETURNING *
		`, params["id"], params["item_id"])

		json.NewDecoder(request.Body).Decode(listItem)

		_, err := infra.UpdateListItem(
			sql,
			listItem.Title, listItem.Description, listItem.UserId, listItem.ListId,
		)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprintf(": %v", err),
			})

			return
		}

		json.NewEncoder(response).Encode(listItem)
	}
}

func (listItem *ListItem) DeleteListItem() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		sql := fmt.Sprintf(
			`DELETE FROM list_item_schema 	
			WHERE list_id=%v AND list_item_id=%v`,

			params["id"], params["item_id"])

		_, err := config.Delete(sql)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"Fail": fmt.Sprintf("Error when deleting user: %v", err),
			})

			return
		}

		json.NewEncoder(response).Encode(map[string]string{
			"message": "Item deleted with success!",
		})
	}
}
