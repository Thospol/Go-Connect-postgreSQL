package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	connectDB()
}
func connectDB() {
	var err error
	connectionString := "postgres://xifczoul:Vd82MWI68IqD9kKL1cLmSxZinSBuoGnN@baasu.db.elephantsql.com:5432/xifczoul"
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect DB Success!")
}
