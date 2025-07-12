package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./store.db")
	if err != nil {
		log.Fatal(err)
		return err
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS purchases (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        item TEXT NOT NULL,
        amount INTEGER NOT NULL,
        price REAL NOT NULL,
        date TEXT NOT NULL
    );`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
