package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://stefan@localhost/stefan?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
	defer db.Close()

	//

	stmt, err := db.Prepare("insert into Things (Name) values ($1) returning ID")
	if err != nil {
		log.Fatal("Cannot prepare statement: ", err)
	}

	things := []string{"Bed", "Table", "Backpack", "Plate"}

	for _, thing := range things {
		var id int
		if err := stmt.QueryRow(thing).Scan(&id); err != nil {
			log.Fatal("Cannot execute prepared statement: ", err)
		}

		fmt.Printf("%d: %s\n", id, thing)
	}
}
