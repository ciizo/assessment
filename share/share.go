package share

import (
	"database/sql"
	"errors"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

type MockDB struct {
	query        string
	lastInsertID int64
	rowsAffected int64
}

func (m *MockDB) LastInsertId() (int64, error) {
	return m.lastInsertID, nil
}
func (m *MockDB) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.query = query
	return m, nil
}

func (m *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	m.query = query

	return &sql.Row{}
}

func (m *MockDB) Prepare(query string) (*sql.Stmt, error) {

	return nil, errors.New("not implement/support ")
}
