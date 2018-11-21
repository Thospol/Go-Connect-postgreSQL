package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db                                          *sql.DB
	id                                          int
	firstname, lastname, email, tel, insertstmt string
)

func main() {
	connectDB()
	//selectAll(db)
	insertstmt = "INSERT INTO users (firstname, lastname, email,tel) values ($1, $2,$3,$4)"
	firstname = "aekawan"
	lastname = "krubsun"
	email = "aekawan@gmail.com"
	tel = "0897979696"
	insertData(db, firstname, lastname, email, tel, insertstmt)

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

func selectAll(db *sql.DB) {
	querystmt := "select * from Users"
	rows, err := db.Query(querystmt)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &firstname, &lastname, &email, &tel)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("{\nID:", id, "\nFirstName:", firstname, "\nLastName:", lastname, "\nEmail:", email, "\nTel:", tel, "\n}")
	}
}

func selectRow(db *sql.DB, idUser int) {
	querystmt := "select * from Users where id = $1"
	row := db.QueryRow(querystmt)
	err := row.Scan(&id, &lastname, &firstname, &email, &tel)
	if err != nil {
		log.Fatal(err)
	}
}

func insertData(db *sql.DB, firstname, lastname, email, tel, insertstmt string) {
	result, err := db.Exec(insertstmt, firstname, lastname, email, tel)
	if err != nil {
		log.Fatal(err)
	}
	numberOfInsert, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result Number Row of Insert:", numberOfInsert)
	selectAll(db)
}

func deleteData(db *sql.DB, id int) {
	deleteStmt := "delete from users where id = $1"
	result, err := db.Exec(deleteStmt, id)

	if err != nil {
		log.Fatal(err)
	}
	numberOfDeleteRow, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result Number Row of Delete:", numberOfDeleteRow)
	selectAll(db)
}
