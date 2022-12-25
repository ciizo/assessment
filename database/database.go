package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func InitDb() {
	db := getDatabase()
	db.initExpenseTable()
}

func NewDb(filename string) *Db {
	return getDatabase()
}

func getDatabase() *Db {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	return &Db{conn}
}

func (db *Db) initExpenseTable() {
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
