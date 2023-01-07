//go:build unit || service
// +build unit service

package expense

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ciizo/assessment/model"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSuccess(t *testing.T) {

	mockDb, mockSql, err := sqlmock.New()
	assert.NoError(t, err)
	id := 1
	entity := &model.Expense{
		Title:  "test-title",
		Amount: 100,
		Note:   "test-note",
		Tags:   []string{"test-tag1", "test-tag2"}}
	mockSql.ExpectPrepare(regexp.QuoteMeta("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1")).ExpectExec().WithArgs(id, entity.Title, entity.Amount, entity.Note, pq.Array(entity.Tags)).WillReturnResult(sqlmock.NewResult(int64(entity.ID), 1))
	setUpByDB(mockDb)

	err = service.Update(id, entity)

	assert.NoError(t, err)

}
