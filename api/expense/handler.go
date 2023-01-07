package expense

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ciizo/assessment/database"
	"github.com/ciizo/assessment/share"
	"github.com/go-playground/validator/v10"

	"github.com/ciizo/assessment/service/expense"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gommon_log "github.com/labstack/gommon/log"
)

type Handler struct {
	expenseService *expense.ExpenseService
}

func setupForTest(t *testing.T) (*httptest.Server, func()) {
	t.Helper()

	share.Validate = validator.New()

	httpHandler := echo.New()
	httpHandler.Logger.SetLevel(gommon_log.INFO)

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
	httpHandler.Logger.SetLevel(gommon_log.INFO)

	registerHandlerByDBForTest(httpHandler, mockDb)

	server := httptest.NewServer(httpHandler)

	teardown := func() {
		server.Close()
	}

	return server, teardown
}

func setupForITTest(t *testing.T) *echo.Echo {
	t.Helper()

	database.InitDb(share.IT_Test_DB_ConnectionString)
	share.Validate = validator.New()

	httpHandler := echo.New()
	httpHandler.Logger.SetLevel(gommon_log.INFO)

	go func(e *echo.Echo) {
		RegisterHandler(e, share.IT_Test_DB_ConnectionString)

		// Start server
		if err := e.Start(fmt.Sprintf(":%d", share.IT_Test_ServerPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err, " shutting down the server")
		}
	}(httpHandler)

	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%v", share.IT_Test_ServerPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}

	return httpHandler
}

func RegisterHandler(httpHandler *echo.Echo, dbConnectionString string) {

	db := database.GetDatabase(dbConnectionString)
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
