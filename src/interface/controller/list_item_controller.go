package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
	"github.com/jhamiltonjunior/priza-tech-backend/src/infra"
)

// Here will contain the items that were inside the "parent" List
// That is, Every item belongs to a list,
// But not every list will have an item, at least it's not mandatory
//
// So there can be empty lists
//
// I mean: /api/v{1}
// 
// Aqui irá conter os itens que ficaram dentro da Lista "pai"
// Ou seja, Todo item pertence a uma lista,
// Mas nem toda lista terá um item, pelo menos não é obrigatório
// 
// Então pode haver listas vazias
//
// Refiro-me: /api/v{1}
// 
type ListItem struct {
	ID          int    `json:"list_item_id" db:"list_item_id"`
	UserId      int    `json:"user_id" db:"user_id"`
	ListId      int    `json:"list_id" db:"list_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

// This function is responsible for creating the list items
// 
// Essa função é responsavel por criar os itens da lista
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

// Will show all items in the list
// the WHERE here refers to the list_id, there is no lookup to see
// only one item in a list, there was no such requirement, maybe that wasn't the point
func (listItem *ListItem) ShowListItem() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		items := []ListItem{}

		sql := fmt.Sprintf("SELECT * FROM list_item_schema WHERE list_id=%v", params["id"])

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

// Will update an item in a list.
//
// You need to pass item id and list id.
//
// In the code below you will see:
//  params["id"], params["item_id"]
//  Where params["id"] refers to list_id, that is, the id of the list,
//  And params["item_id"] refers to list_item_id, that is, the id of the item in the list
// 
// Portuguese Version
// Vai fazer o update de uma item de uma lista.
// 
// Você precisa passar o id do item e o id da lista.
// 
// No codigo abaixo você vai ver:
//  params["id"], params["item_id"]
//  Onde params["id"] é referente a list_id, ou seja, o id da lista,
//  E params["item_id"] é refereinte a list_item_id, ou seja, o id do item que está na lista
// 
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

// This function is responsible for deleting an item from the list.
//
// This function is a little different from the others.
// It won't just delete using the list_item_id, it also needs the list_id
//
// That is, I will delete an item referring to the parent list, I would not need to pass the id
// from the parent list, since there is only one id referring to the list_item_id.
//
// But maybe this is a safer way to delete an item from the list
// 
// Portuguese Version
// Esta função é responsavel por deletar um item da lista.
// 
// Ela é um pouco diferente das outras.
// Ela não vai deletar apenas usando o list_item_id, ela pricisa também do list_id
// 
// Ou seja, eu vou deletar um item referente a lista pai, eu não precisaria passar o id
// da lista pai, já que só exite um id referente ao list_item_id.
// 
// Mas talvez essa seja uma forma mais segura de deletar um item da lista
// 
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
