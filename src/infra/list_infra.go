package infra

import (
	"database/sql"
	"os"

	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
)
// checked bool,
func InsertList(sql, title string,  userId int) (sql.Result, error) {
	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(sql, title, userId)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func UpdateList(sql, title string, checked bool) (sql.Result, error) {
	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(sql, title, checked)
	if err != nil {
		return nil, err
	}

	return result, nil
}

