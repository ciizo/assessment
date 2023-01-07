//go:build integration
// +build integration

package expense

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"strings"
	"testing"
	"time"

	"github.com/ciizo/assessment/model"
	"github.com/ciizo/assessment/share"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const authKeyScheme = "November"

var authKey = fmt.Sprintf("%s 10, 2009", authKeyScheme)

func TestITGetGreeting(t *testing.T) {

	// Setup server
	eh := setupForITTest(t)

	// Arrange
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/", share.IT_Test_ServerPort), strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authKey)
	client := http.Client{}

	// Act
	resp, err := client.Do(req)
	assert.NoError(t, err)
	resp.Body.Close()

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}

func TestITCreateUpdateListRead(t *testing.T) {
	// Setup server
	eh := setupForITTest(t)

	entity := &model.Expense{
		Title:  "it-test-title",
		Amount: 100,
		Note:   "it-test-note",
		Tags:   []string{"it-test-tag1", "it-test-tag2"}}

	t.Run("Create expense", func(t *testing.T) {

		// Arrange
		entityBytes, err := json.Marshal(entity)
		assert.NoError(t, err)
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/expenses", share.IT_Test_ServerPort), bytes.NewReader(entityBytes))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		req.Header.Set(echo.HeaderAuthorization, authKey)
		client := http.Client{}

		// Act
		resp, err := client.Do(req)
		assert.NoError(t, err)

		byteBody, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		resp.Body.Close()

		// Assertions
		actual := &model.Expense{}
		json.Unmarshal(byteBody, actual)
		entity.ID = actual.ID

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, resp.StatusCode)
			assert.Greater(t, actual.ID, 0)
			assert.Equal(t, entity, actual)
		}

	})

	t.Run("Update expense", func(t *testing.T) {

		// Arrange
		entity.Amount = 300
		entity.Title = entity.Title + "updated"
		entityBytes, err := json.Marshal(entity)
		assert.NoError(t, err)
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:%d/expenses/%v", share.IT_Test_ServerPort, entity.ID), bytes.NewReader(entityBytes))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		req.Header.Set(echo.HeaderAuthorization, authKey)
		client := http.Client{}

		// Act
		resp, err := client.Do(req)
		assert.NoError(t, err)

		byteBody, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		resp.Body.Close()

		// Assertions
		actual := &model.Expense{}
		json.Unmarshal(byteBody, actual)
		entity.ID = actual.ID

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, entity, actual)
		}

	})

	t.Run("Get list of expense", func(t *testing.T) {

		// Arrange
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses", share.IT_Test_ServerPort), strings.NewReader(""))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		req.Header.Set(echo.HeaderAuthorization, authKey)
		client := http.Client{}

		// Act
		resp, err := client.Do(req)
		assert.NoError(t, err)

		byteBody, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		resp.Body.Close()

		// Assertions
		actual := []model.Expense{}
		json.Unmarshal(byteBody, &actual)

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Greater(t, len(actual), 0)
		}

	})

	t.Run("Get expense", func(t *testing.T) {

		// Arrange
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/expenses/%v", share.IT_Test_ServerPort, entity.ID), strings.NewReader(""))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		req.Header.Set(echo.HeaderAuthorization, authKey)
		client := http.Client{}

		// Act
		resp, err := client.Do(req)
		assert.NoError(t, err)

		byteBody, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		resp.Body.Close()

		// Assertions
		actual := &model.Expense{}
		json.Unmarshal(byteBody, actual)
		entity.ID = actual.ID

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, entity, actual)
		}

	})

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := eh.Shutdown(ctx)
	assert.NoError(t, err)
}
