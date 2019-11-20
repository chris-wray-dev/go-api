package main

import (
	"database/sql"
	"fmt"

	// Import log for logging errors, net/http as we need to make api requests.
	// We install Gorilla Mux so we can handle path and query parameters. 'net/http' does not allow us to do this.
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "project_test"
)

// This is where we would establish our connection to the PSQL database
// Below is where we query the database:
func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d "+
		"dbname=%s sslmode=disable",
		host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlStatement := `SELECT * FROM locations WHERE location_id=$1;` // the SQL query statement
	var id int
	var name string
	var text string                                  // The variable name we use to store result of query
	row := db.QueryRow(sqlStatement, 3)              // We pass the db query to our DB, along with the ID we are looking for. Info is return to 'row'.
	switch err := row.Scan(&id, &name, &text); err { // Here, the returned info in row is given to the variable 'name'
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, name, text)
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
