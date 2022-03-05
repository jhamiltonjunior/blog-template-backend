package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Open(dbsourse string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("postgres", dbsourse)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("erro no ping, tio: %v", err)
		return
	}

	return
}

func Insert(sql string, values []string) {

}

func Select(sql string) (*sqlx.DB, error) {

	db, err := Open(
		// "postgres://postgres@localhost/vibbra?sslmode=disable",
		"postgres://postgres:0000@localhost/vibbra?sslmode=disable",
	)
	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	_, err = db.Exec(sql)
	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	return db, nil
}
