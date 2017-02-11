package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func deleteAllThings(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Print("Cannot begin transaction: ", err)
		return
	}

	if _, err := tx.Exec("delete from Things"); err != nil {
		log.Print("Cannnot exec: ", err)
		return
	}

	if err := tx.Rollback(); err != nil {
		log.Print("Cannot rollback: ", err)
		return
	}
}

func main() {
	db, err := sql.Open("postgres", "postgres://stefan@localhost/stefan?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
	defer db.Close()

	//

	deleteAllThings(db)
}
