//go:build unit || service
// +build unit service

package expense

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ciizo/assessment/model"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetSuccess(t *testing.T) {

	mockDb, mockSql, err := sqlmock.New()
	assert.NoError(t, err)
	entity := &model.Expense{
		ID:     1,
		Title:  "test-title",
		Amount: 100,
		Note:   "test-note",
		Tags:   []string{"test-tag1", "test-tag2"}}
	newsMockRows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(entity.ID, entity.Title, entity.Amount, entity.Note, pq.Array(entity.Tags))
	mockSql.ExpectPrepare("SELECT id, title, amount, note, tags FROM expenses").ExpectQuery().WithArgs(entity.ID).WillReturnRows(newsMockRows)
	setUpTestServiceByDB(mockDb)

	_, err = service.Get(entity.ID)

	assert.NoError(t, err)

}
