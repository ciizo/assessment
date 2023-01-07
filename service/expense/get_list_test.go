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

func TestGetListSuccess(t *testing.T) {

	mockDb, mockSql, err := sqlmock.New()
	assert.NoError(t, err)
	entity := &model.Expense{
		Title:  "test-title",
		Amount: 100,
		Note:   "test-note",
		Tags:   []string{"test-tag1", "test-tag2"}}
	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(1, entity.Title, entity.Amount, entity.Note, pq.Array(entity.Tags)).
		AddRow(2, entity.Title, entity.Amount, entity.Note, pq.Array(entity.Tags))
	mockSql.ExpectPrepare(regexp.QuoteMeta("SELECT id, title, amount, note, tags FROM expenses")).ExpectQuery().WillReturnRows(newsMockRows)
	setUpByDB(mockDb)

	_, err = service.GetList()

	assert.NoError(t, err)

}
