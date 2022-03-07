package infra

import (
	"database/sql"
	"os"
	"time"

	"github.com/jhamiltonjunior/priza-tech-backend/src/config"
)


func InsertUser(sql, username, fullname, email, passwd string) (sql.Result, error) {
	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(sql, username, fullname, email, passwd)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func SelectUser(sql string) (*sql.Rows, error) {

	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)
	if err != nil {
		return nil, err
	}

	result, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateUser(sql, username, fullname, email, passwd string, updateAt time.Time) (sql.Result, error) {
	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(sql, username, fullname, email, passwd, updateAt)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteUser(sql string) (sql.Result, error) {
	db, err := config.Open(
		os.Getenv("DB_SOURCE"),
	)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(sql)
	if err != nil {
		return nil, err
	}

	return result, nil
}
