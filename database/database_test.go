package database

import (
	"database/sql"
	"testing"

	"github.com/ciizo/assessment/model"
)

type mockDB struct {
	query        string
	lastInsertID int64
	rowsAffected int64
}

func (m *mockDB) LastInsertId() (int64, error) {
	return m.lastInsertID, nil
}
func (m *mockDB) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.query = query
	return m, nil
}

func (m *mockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	m.query = query

	// return sql.Row(&mockSqlROW{})
	// return (*sql.Row)(unsafe.Pointer(&mockSqlROW{}))
	return &mockSqlROW{}
}

type mockSqlROW struct {
	err  error
	rows *sql.Rows
}

func (r *mockSqlROW) Scan(dest ...interface{}) error {
	// if r.err != nil {
	// 	return r.err
	// }

	return nil
}

func (r *mockSqlROW) Err() error {
	return r.err
}

func TestCreateExpense(t *testing.T) {
	mock := &mockDB{}
	db := &Db{mock}
	entity := &model.Expense{Title: "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"}}

	err := db.CreateExpense(entity)

	if err != nil {
		t.Error(err)
	}

}
