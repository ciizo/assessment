package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ciizo/assessment/model"
	"github.com/lib/pq"
)

type DB interface {
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Db struct {
	DB
	IsTestMode bool
}

func InitDb() {
	db := getDatabase()
	db.initExpenseTable()
}

func NewDb() *Db {
	return getDatabase()
}

func getDatabase() *Db {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	return &Db{DB: conn}
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

func (db *Db) CreateExpense(entity *model.Expense) error {

	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id",
		entity.Title, entity.Amount, entity.Note, pq.Array(&entity.Tags))

	var err error
	if !db.IsTestMode {
		err = row.Scan(&entity.ID)
	}
	if err != nil {
		fmt.Println("can't create expense ", err)
		return err
	}

	return nil
}
