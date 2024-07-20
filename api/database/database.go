package database

import (
	"database/sql"
	"github.com/ch3yb/clinic/env"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func StartDatabase() (*sql.DB, error) {

	dbFilePath := env.Conf.Database.URL
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Println("err when Open Database: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("err database.go: ", err)
		return nil, err
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+env.Conf.Database.FILE, "sqlite", driver)

	if err != nil {
		panic(err.Error() + "he")
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println("here")
		log.Fatal(err)
	}
	return db, nil
}
