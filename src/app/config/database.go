package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
}

func CloseDB() {
	db.Close()
	fmt.Println("Database connection closed")
}
