package learn_go_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	/* username:password@tcp(host:port)/dbname?param=value */
	db, err := sql.Open("mysql", "root:Guntur123!@tcp(localhost:3306)/learn_go_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// using db
}
