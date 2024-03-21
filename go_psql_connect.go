package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "0.0.0.0"
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func getTablesNames(db *sql.DB) {
	// some queries:
	stmt := "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'"
	rows, err := db.Query(stmt)

	CheckError(err)
	defer rows.Close()
	// print all public tables names
	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			panic(err)
		}
		fmt.Println(tableName)
	}
}

func getAllUsers(db *sql.DB) {
	stmt := "SELECT * FROM users"
	rows, err := db.Query(stmt)

	CheckError(err)
	defer rows.Close()
	// print all usersfrom the table
	var userid, username, userpass string
	for rows.Next() {
		if err := rows.Scan(&userid, &username, &userpass); err != nil {
			panic(err)
		}
		fmt.Println("User ID:", userid)
		fmt.Println("Username:", username)
		fmt.Println("Password:", userpass)
		fmt.Println("#####")
	}
}

func connect() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	fmt.Println("Connected!")

	// close database
	defer db.Close()

	// getTablesNames(db)
	getAllUsers(db)

	// check db
	err = db.Ping()
	CheckError(err)

}
