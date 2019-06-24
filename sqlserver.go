package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gcmurphy/getpass"
)

var db *sql.DB

var server = "vivnetwork.bi.exigo.com"
var port = 1433
var user = "VivNetwork_Kevin"
var password string
var database = "VivNetworkReporting"

func main() {
	password, _ = getpass.GetPass()
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	tsql := fmt.Sprintf("SELECT top 5 CustomerID, FirstName, LastName from Customers;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var fname, lname string
		var id int

		// Get values from row.
		err := rows.Scan(&id, &fname, &lname)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("ID: %d, FirstName: %s, LastName: %s\n", id, fname, lname)
		count++
	}

}
