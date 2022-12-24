package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	db = getDatabase()
	initExpenseTable()
}

func getDatabase() *sql.DB {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	return conn
}

func initExpenseTable() {
	createTableSql := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`

	if _, err := db.Exec(createTableSql); err != nil {
		log.Fatal("can't create table ", err)
	}
}
