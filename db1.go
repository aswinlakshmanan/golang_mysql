package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:hanuman1998@tcp(localhost:3306)/database1")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmIns, err := db.Prepare("INSERT INTO squarenum VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmIns.Close()

	stmOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE NUMBER = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmOut.Close()

	for i := 0; i < 25; i++ {
		_, err = stmIns.Exec(i, (i * i))
		if err != nil {
			panic(err.Error())
		}
	}
	var squareNum int
	err = stmOut.QueryRow(13).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)
	err = stmOut.QueryRow(1).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)
}
