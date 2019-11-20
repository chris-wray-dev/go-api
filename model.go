package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "cillian"
	password = "liverpool2"
	dbname   = "project_test"
)

// This is where we would establish our connection to the PSQL database
// Below is where we query the database:
func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	sqlStatement := `SELECT location_name FROM locations WHERE location_id=$1;` // the SQL query statement
	var name string                                                             // The varialbe name we use to store result of query
	row := db.QueryRow(sqlStatement, 3)                                         // We pass the db query to our DB, along with the ID we are looking for. Info is return to 'row'.
	switch err := row.Scan(&name); err {                                        // Here, the returned info in row is given to the variable 'name'
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(name)
	default:
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
		// results, err:= db.Query("SELECT location_name FROM locations")
	}
	fmt.Println("Successfully connected!")
}
