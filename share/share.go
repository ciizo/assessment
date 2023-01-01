package share

import (
	"database/sql"

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

	// return sql.Row(&mockSqlROW{})
	// return (*sql.Row)(unsafe.Pointer(&mockSqlROW{}))
	// return &mockSqlROW{}
	return &sql.Row{}
}

// type mockSqlROW struct {
// 	err  error
// 	rows *sql.Rows
// }

// func (r *mockSqlROW) Scan(dest ...interface{}) error {
// 	// if r.err != nil {
// 	// 	return r.err
// 	// }

// 	return nil
// }

// func (r *mockSqlROW) Err() error {
// 	return r.err
// }
