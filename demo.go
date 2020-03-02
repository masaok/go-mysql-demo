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

// User ...
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// https://tutorialedge.net/golang/golang-mysql-tutorial/
	// https://www.geeksforgeeks.org/different-ways-to-concatenate-two-strings-in-golang

	// Open up our database connection.
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// https://tutorialedge.net/golang/golang-mysql-tutorial/#populating-structs-from-results

	// Execute the query
	results, err := db.Query("SELECT user_id, username FROM `user`")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var user User
		// for each row, scan the result into our user composite object
		err = results.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the user's Name attribute
		log.Printf(user.Name)

		// https://www.golangprograms.com/how-to-print-struct-variables-data.html
		log.Printf("%+v\n", user)
	}

}
