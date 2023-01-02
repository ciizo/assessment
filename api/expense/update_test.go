package expense

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ciizo/assessment/model"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestHttpUpdate(t *testing.T) {
	mockDb, mockSql, err := sqlmock.New()
	assert.NoError(t, err)

	id := 1
	entity := &model.Expense{
		Title:  "test-title",
		Amount: 100,
		Note:   "test-note",
		Tags:   []string{"test-tag1", "test-tag2"}}
	mockSql.ExpectPrepare(regexp.QuoteMeta("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1")).ExpectExec().WithArgs(id, entity.Title, entity.Amount, entity.Note, pq.Array(entity.Tags)).WillReturnResult(sqlmock.NewResult(int64(entity.ID), 1))

	entityBytes, err := json.Marshal(entity)
	assert.NoError(t, err)

	server, teardown := setupByDBForTest(t, mockDb)
	defer teardown()

	t.Run("Update expense success", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/expenses/%v", server.URL, id), bytes.NewReader(entityBytes))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := http.DefaultClient.Do(req)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.NotNil(t, resp.Body)
		}

		actual := &model.Expense{}
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)

		json.Unmarshal(body, actual)
		entity.ID = id
		assert.Equal(t, entity, actual)

	})
}
