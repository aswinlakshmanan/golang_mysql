package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("GO MYSQL TUTORIAL")

	db, err := sql.Open("mysql", "root:hanuman1998@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO users VALUES('ELIOT')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	fmt.Println("succesfully inserted into users table.")
}
