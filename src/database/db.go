package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDb(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		panic(err)
	}

	return db
}
