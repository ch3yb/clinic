package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func StartDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./clinic.db")
	if err != nil {
		log.Println("err when Open Database: ", err)
		return nil, err
	}
	defer db.Close()
	return db, nil
}
