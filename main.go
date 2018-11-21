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
	tel = "0897979696"
	updateData(db, tel, 4)

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
	querystmt := "select * from Users Order by id asc"
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

func insertData(db *sql.DB, firstname, lastname, email, tel string) {
	insertstmt = "INSERT INTO users (firstname, lastname, email,tel) values ($1, $2,$3,$4)"
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

func updateData(db *sql.DB, tel string, id int) {
	updateStmt := "update users set tel = $1 where id = $2"
	result, err := db.Exec(updateStmt, tel, id)
	if err != nil {
		log.Fatal(err)
	}
	numberOfUpdateRow, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result Number Row of updateData:", numberOfUpdateRow)
	selectAll(db)
}
