package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetData() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "shop.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
