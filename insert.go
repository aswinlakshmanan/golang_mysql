package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type post struct {
	Id   int
	Name string
	Age  int
}

func main() {
	db, err := sql.Open("mysql", "root:hanuman1998@tcp(localhost:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into users(id,name,age) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("5", "five", "33")
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert id", id)

}
