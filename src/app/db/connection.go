package db

import (
	"database/sql"

	"app/config"
)

var db *sql.DB

func InitDB() error {
	var err error

	db, err = config.GetDB()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
