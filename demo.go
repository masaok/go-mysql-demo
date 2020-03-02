package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// https://gobyexample.com/environment-variables
var dbhost string = os.Getenv("DBHOST")
var dbuser string = os.Getenv("DBUSER")
var dbpass string = os.Getenv("DBPASS")
var dbname string = os.Getenv("DBNAME")

func main() {
	fmt.Println("Go MySQL Tutorial")

	// https://tutorialedge.net/golang/golang-mysql-tutorial/

	// Open up our database connection.
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM `user`")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}

}
