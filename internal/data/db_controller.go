package data

import (
	"database/sql"
	"fmt"

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


func InsertLink(db *sql.DB, link *Link) (value *Link, error *error) {

	_, err := db.Exec("INSERT INTO nips (id, url) VALUES (?, ?)", fmt.Sprintf("%+v", link.Id), link.Url)
	if err != nil {
		println("there is an error", err.Error())
		return nil, &err
		// panic(err.Error())
	} else {
		println("there is no error")
		return link, nil
	}

} 


func GetLink(db *sql.DB, id uint64) *Link {

	var link Link

	result, err := db.Query("SELECT id, url FROM nips WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		error := result.Scan(&link.Id, &link.Url)
		if error != nil {
			panic(error.Error())
		}
	}

	return &link
}