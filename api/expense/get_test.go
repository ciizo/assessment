package expense

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ciizo/assessment/model"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestHttpGet(t *testing.T) {
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

	server, teardown := setupByDBForTest(t, mockDb)
	defer teardown()

	t.Run("Get expense success", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/expenses/%v", server.URL, entity.ID))
		if err != nil {
			t.Fatalf("error http Get %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("statusCode expected %d, got %d", http.StatusOK, resp.StatusCode)
		}

		actual := &model.Expense{}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("error ioutil read resp.Body: %v", err)
		}

		json.Unmarshal(body, actual)
		if !reflect.DeepEqual(actual, entity) {
			t.Errorf("expected (%v), got (%v)", entity, actual)
		}

	})
}
