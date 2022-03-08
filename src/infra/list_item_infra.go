package infra

import (
	"database/sql"
	"os"

	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
	"github.com/jmoiron/sqlx"
)

// Eu estou usando essa função aqui apesar dela não ser tão necessária,
//  Mas isso evita deixar o list_item_controller.go verboso
func SelectListItem(sql string) (*sqlx.DB, error) {

	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertListItem(sql string, userId, listId int, title, description string) (sql.Result, error) {
	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(sql, userId, listId, title, description)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func UpdateListItem(sql, title, description string, userId, listId int) (sql.Result, error) {
	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)

	if err != nil {
		return nil, err
	}

	result, err := db.Exec(sql, title, description, userId, listId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
