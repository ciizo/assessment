package expense

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/ciizo/assessment/database"
	"github.com/ciizo/assessment/share"
	"github.com/go-playground/validator/v10"

	"github.com/ciizo/assessment/service/expense"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	expenseService *expense.ExpenseService
}

func setupForTest(t *testing.T) (*httptest.Server, func()) {
	t.Helper()

	share.Validate = validator.New()

	httpHandler := echo.New()
	registerHandlerForTest(httpHandler)

	server := httptest.NewServer(httpHandler)

	teardown := func() {
		server.Close()
	}

	return server, teardown
}

func setupByDBForTest(t *testing.T, mockDb *sql.DB) (*httptest.Server, func()) {
	t.Helper()

	share.Validate = validator.New()

	httpHandler := echo.New()
	registerHandlerByDBForTest(httpHandler, mockDb)

	server := httptest.NewServer(httpHandler)

	teardown := func() {
		server.Close()
	}

	return server, teardown
}

func RegisterHandler(httpHandler *echo.Echo) {

	db := database.NewDb()
	service := expense.NewService(db)
	handler := &Handler{expenseService: service}

	httpHandler.Use(middleware.Logger())
	httpHandler.Use(middleware.Recover())

	registerRoutes(httpHandler, handler)
}

func registerHandlerForTest(httpHandler *echo.Echo) {

	mock := &share.MockDB{}
	db := &database.Db{DB: mock, IsTestMode: true}
	service := expense.NewService(db)
	handler := &Handler{expenseService: service}

	httpHandler.Use(middleware.Recover())

	registerRoutes(httpHandler, handler)
}

func registerHandlerByDBForTest(httpHandler *echo.Echo, mockDb *sql.DB) {

	db := &database.Db{DB: mockDb, IsTestMode: true}
	service := expense.NewService(db)
	handler := &Handler{expenseService: service}

	httpHandler.Use(middleware.Recover())

	registerRoutes(httpHandler, handler)
}

func registerRoutes(httpHandler *echo.Echo, handler *Handler) {

	httpHandler.POST("/expenses", handler.createHandler)
	httpHandler.GET("/expenses/:id", handler.getHandler)
}
