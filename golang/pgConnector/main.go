package main

import(
        "database/sql"
        "fmt"
        _ "github.com/lib/pq"
)


func main () {
	pgInfo := "user=postgres password=docker dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres",pgInfo)

	if err != nil {
		panic(err)
	}
	
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
