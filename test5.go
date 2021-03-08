package main

import (
	// "database/sql"
	"log"
	// _ "github.com/go-sql-driver/mysql"
)

type Rep struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// // Open up our database connection.
	// db, err := sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/testdb")

	// // if there is an error opening the connection, handle it
	// if err != nil {
	// 	log.Print(err.Error())
	// }
	// defer db.Close()

	// // Execute the query
	// results, err := db.Query("SELECT id, name FROM Reps")
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// for results.Next() {
	// 	var rep Rep
	// 	// for each row, scan the result into our tag composite object
	// 	err = results.Scan(&rep.ID, &rep.Name)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	// 	}
	// 	// and then print out the tag's Name attribute
	// 	log.Printf(rep.Name)
	// }
	log.Printf("OK")

}
