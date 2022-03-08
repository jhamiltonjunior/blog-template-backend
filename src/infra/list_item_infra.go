package infra

import (
	"database/sql"
	"os"

	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
	"github.com/jmoiron/sqlx"
)

func SelectListItem(sql string) (*sqlx.Rows, error) {

	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)
	if err != nil {
		return nil, err
	}

	result, err := db.Queryx(sql)
	if err != nil {
		return nil, err
	}

	return result, nil
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

func UpdateListItem(sql, title, description string) (sql.Result, error) {
	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)

	if err != nil {
		return nil, err
	}

	result, err := db.Exec(sql, title, description)
	if err != nil {
		return nil, err
	}

	return result, nil
}
