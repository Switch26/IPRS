package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const dbFileLocation string = "lookup_database.sqlite3"

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", dbFileLocation)
	CheckError(err)
	return db
}

func createDB() {
	db := connectDB()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error: ", r)
		}
	}()
	// schema
	const createCommand string = `
		CREATE TABLE IF NOT EXISTS lookup_table (
		cid TEXT,
		url TEXT
		);`
	result, err := db.Exec(createCommand)
	CheckError(err)
	fmt.Println("Did SQLite create table:", result)
}

func insertRow(cid, url string) (sql.Result, error) {
	db := connectDB()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error: ", r)
		}
	}()

	result, err := db.Exec("INSERT INTO lookup_table VALUES (?,?);", cid, url)
	return result, err
}
