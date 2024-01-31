package learn_go_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:Guntur123!@tcp(localhost:3306)/learn_go_database")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
