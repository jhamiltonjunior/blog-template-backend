package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Open(dbsourse string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("postgres", dbsourse)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	return
}
