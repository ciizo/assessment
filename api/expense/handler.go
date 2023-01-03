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
	"github.com/labstack/gommon/log"
)

type Handler struct {
	expenseService *expense.ExpenseService
}

func setupForTest(t *testing.T) (*httptest.Server, func()) {
	t.Helper()

	share.Validate = validator.New()

	httpHandler := echo.New()
	httpHandler.Logger.SetLevel(log.INFO)

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
	httpHandler.Logger.SetLevel(log.INFO)

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
	httpHandler.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator:  authMiddleware,
		KeyLookup:  "header:" + echo.HeaderAuthorization,
		AuthScheme: "November",
	}))

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
	httpHandler.PUT("/expenses/:id", handler.updateHandler)
	httpHandler.GET("/expenses", handler.getListHandler)

}

func authMiddleware(auth string, c echo.Context) (bool, error) {
	if "10, 2009" == auth {
		return true, nil
	}

	return false, nil

}
