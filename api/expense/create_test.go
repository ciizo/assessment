package expense

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/ciizo/assessment/model"
)

func TestHttpCreate(t *testing.T) {
	server, teardown := setupForTest(t)
	defer teardown()

	t.Run("Create expense success", func(t *testing.T) {

		want := &model.Expense{
			Title:  "Test T",
			Amount: 1,
			Note:   "Test N",
			Tags:   []string{"T1", "T2"}}

		entityBytes, err := json.Marshal(want)
		if err != nil {
			t.Fatalf("error Marshal %v", err)
		}

		resp, err := http.Post(fmt.Sprintf("%s/expenses", server.URL), "application/json", bytes.NewReader(entityBytes))
		if err != nil {
			t.Fatalf("error http SET %v", err)
		}

		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("statusCode expected %d, got %d", http.StatusCreated, resp.StatusCode)
		}

		actual := &model.Expense{}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("error ioutil read resp.Body: %v", err)
		}

		json.Unmarshal(body, actual)
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("expected (%v), got (%v)", want, actual)
		}

	})
}
