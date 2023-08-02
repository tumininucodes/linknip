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


func GetLink(db *sql.DB, id uint64) *Link {

	var link Link

	result, err := db.Query("SELECT id, url FROM nips WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	// for result.Next() {

	// }
	error := result.Scan(&link.Id, &link.Url)
	if error != nil {
		panic(error.Error())
	}

	return &link
}