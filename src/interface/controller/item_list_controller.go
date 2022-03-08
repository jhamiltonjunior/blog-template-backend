package controller

import "net/http"

type ListItem struct {
	ID          int    `json:"list_id" db:"list_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UserId      int    `json:"user_id" db:"user_id"`
}

func (listItem *ListItem) CreateListItem() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

	}
}

func (listItem *ListItem) ShowListItem() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

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
