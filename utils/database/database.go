package database

import (
	"database/sql"
)

var db *sql.DB

func Init(database *sql.DB) {
	db = database
}

func GetDB() *sql.DB {
	return db
}
