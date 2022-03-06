package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Open(dbsourse string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("postgres", dbsourse)
	if err != nil {
		err = fmt.Errorf("erro in open database: %v", err)
		return
	}

	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("erro in ping database: %v", err)
		return
	}

	return
}

func Insert(sql string, values []string) {

}

func Select(sql string) (*sql.Rows, error) {

	db, err := Open(
		os.Getenv("DB_SOURCE"),
	)
	if err != nil {
		return nil, err
	}

	result, err := db.Query(sql)
	if err != nil {
		err = fmt.Errorf("erro in query database: %v", err)
		return nil, err
	}

	return result, nil
}
