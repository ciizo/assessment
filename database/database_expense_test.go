//go:build unit || database
// +build unit database

package database

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ciizo/assessment/model"
	"github.com/ciizo/assessment/share"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpenseSuccess(t *testing.T) {
	mock := &share.MockDB{}
	db := &Db{DB: mock, IsTestMode: true}
	entity := &model.Expense{Title: "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"}}

	err := db.CreateExpense(entity)

	if err != nil {
		t.Error(err)
	}

}

func TestGetExpenseSuccess(t *testing.T) {
	entity := &model.Expense{
		ID:     1,
		Title:  "test-title",
		Amount: 100,
		Note:   "test-note",
		Tags:   []string{"test-tag1", "test-tag2"}}
	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(entity.ID, entity.Title, entity.Amount, entity.Note, pq.Array(entity.Tags))
	mockDb, mockSql, err := sqlmock.New()
	assert.NoError(t, err)

	mockSql.ExpectPrepare(regexp.QuoteMeta("SELECT id, title, amount, note, tags FROM expenses WHERE id=$1")).ExpectQuery().WithArgs(entity.ID).WillReturnRows(newsMockRows)
	db := &Db{DB: mockDb, IsTestMode: true}

	result := &model.Expense{}
	result, err = db.GetExpense(entity.ID)

	if assert.NoError(t, err) {
		assert.NotNil(t, result)
		assert.Equal(t, entity, result)
	}

}

func TestUpdateExpenseSuccess(t *testing.T) {
	entity := &model.Expense{
		ID:     1,
		Title:  "test-title",
		Amount: 100,
		Note:   "test-note",
		Tags:   []string{"test-tag1", "test-tag2"}}
	mockDb, mockSql, err := sqlmock.New()
	assert.NoError(t, err)

	mockSql.ExpectPrepare(regexp.QuoteMeta("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1")).ExpectExec().WithArgs(entity.ID, entity.Title, entity.Amount, entity.Note, pq.Array(entity.Tags)).WillReturnResult(sqlmock.NewResult(int64(entity.ID), 1))
	db := &Db{DB: mockDb, IsTestMode: true}

	err = db.UpdateExpense(entity)

	assert.NoError(t, err)

}

func TestGetExpensesSuccess(t *testing.T) {
	entity := &model.Expense{
		Title:  "test-title",
		Amount: 100,
		Note:   "test-note",
		Tags:   []string{"test-tag1", "test-tag2"}}
	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(1, entity.Title, entity.Amount, entity.Note, pq.Array(entity.Tags)).
		AddRow(2, entity.Title, entity.Amount, entity.Note, pq.Array(entity.Tags))
	mockDb, mockSql, err := sqlmock.New()
	assert.NoError(t, err)

	mockSql.ExpectPrepare(regexp.QuoteMeta("SELECT id, title, amount, note, tags FROM expenses")).ExpectQuery().WillReturnRows(newsMockRows)
	db := &Db{DB: mockDb, IsTestMode: true}

	results := &[]model.Expense{}
	results, err = db.GetExpenses()

	if assert.NoError(t, err) {
		assert.NotNil(t, results)
	}

}
