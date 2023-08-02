package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


func OpenDB() *sql.DB {

	db, error := sql.Open("mysql", "root:alade2001@tcp(localhost:3306)/linknipdb")
	if error != nil {
		panic(error.Error())
	}

	error = db.Ping()

	if error != nil {
		panic(error.Error())
	}

	return db
}


func InsertLink(db *sql.DB, link *Link) *Link {

	_, err := db.Exec("INSERT IGNORE INTO nips (id, url) VALUES (?, ?)", link.Id, link.Url)
	if err != nil {
		panic(err.Error())
	}

	return link
}