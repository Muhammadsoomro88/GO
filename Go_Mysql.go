package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/practice")
	if err != nil {
		fmt.Println("error validating Sql Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Succesfully connected to the database")

	/*err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}*/

	insert, err := db.Query("INSERT INTO accounts(emp_id, salary) VALUES('emp_17',38000)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successful Inserted to Database!")
}
