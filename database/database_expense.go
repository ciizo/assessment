package database

import (
	"fmt"

	"github.com/ciizo/assessment/model"
	"github.com/lib/pq"
)

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

func (db *Db) GetExpense(id int) (*model.Expense, error) {

	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id=$1")
	if err != nil {
		fmt.Println("can'tprepare query one row statment", err)
		return nil, err

	}

	row := stmt.QueryRow(id)
	expense := &model.Expense{}
	err = row.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		fmt.Println("can't Scan row into variables", err)
		return nil, err

	}

	return expense, nil
}

func (db *Db) GetExpenses() (*[]model.Expense, error) {

	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		fmt.Println("can't prepare query rows statment", err)
		return nil, err

	}

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("can't query rows statment", err)
		return nil, err

	}
	expenses := []model.Expense{}
	for rows.Next() {
		expense := model.Expense{}
		err := rows.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
		if err != nil {
			fmt.Println("can't Scan rows into variables", err)
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return &expenses, nil
}

func (db *Db) UpdateExpense(entity *model.Expense) error {
	stmt, err := db.Prepare(`
	UPDATE expenses
	SET title=$2, amount=$3, note=$4, tags=$5
	WHERE id=$1
	`)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = stmt.Exec(entity.ID, entity.Title, entity.Amount, entity.Note, pq.Array(entity.Tags))
	if err != nil {
		fmt.Println("can't update entity ", err)
		return err
	}

	return nil
}
