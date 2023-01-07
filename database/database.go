package database

import (
	"database/sql"
	"log"

	_ "embed"
)

//go:embed script/01-init.sql
var sql_01_init string

type DB interface {
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
}

type Db struct {
	DB
	IsTestMode bool
}

func InitDb(connectionString string) {
	db := GetDatabase(connectionString)
	db.initExpenseTable()
}

func GetDatabase(connectionString string) *Db {
	conn, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	return &Db{DB: conn}
}

func (db *Db) initExpenseTable() {
	if _, err := db.Exec(sql_01_init); err != nil {
		log.Fatal("can't create table ", err)
	}
}
