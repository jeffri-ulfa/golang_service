package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	var err error
	// db, err := sql.Open("mysql", "root:@/kasumi_development")
	db, err := sql.Open("mysql", "godx1:G0LangDX_1@tcp(mysql_lara:3306)/kasumi_dx")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
