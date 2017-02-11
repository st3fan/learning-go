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

	// This is a lot of code for a query

	if true {
		rows, err := db.Query("select Id,Name from Things")
		if err != nil {
			log.Fatal("Cannot query database: ", err)
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				log.Fatal("Cannot scan row: ", err)
			}

			fmt.Printf("%d %s\n", id, name)
		}

		if err := rows.Err(); err != nil {
			log.Fatal("Cannot query rows: ", err)
		}
	}

	// Single row queries are simpler

	if true {
		var name string
		if err := db.QueryRow("select Name from Things where ID = $1", 3).Scan(&name); err != nil {
			log.Fatal("Cannot query single row: ", err)
		}
		fmt.Printf("%s\n", name)
	}

	// Single row queries return the first row even if there are multiple rows

	if true {
		var name string
		if err := db.QueryRow("select Name from Things where ID > $1", 3).Scan(&name); err != nil {
			log.Fatal("Cannot query single row: ", err)
		}
		fmt.Printf("%s\n", name)
	}
}
