package expense

import (
	"github.com/ciizo/assessment/database"
	"github.com/ciizo/assessment/share"

	"github.com/ciizo/assessment/service/expense"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	expenseService *expense.ExpenseService
}

func RegisterHandler(httpHandler *echo.Echo) {

	db := database.NewDb()
	service := expense.NewService(db)
	handler := &Handler{expenseService: service}

	httpHandler.Use(middleware.Logger())
	httpHandler.Use(middleware.Recover())

	registerRoutes(httpHandler, handler)
}

func RegisterHandlerForTest(httpHandler *echo.Echo) {

	mock := &share.MockDB{}
	db := &database.Db{DB: mock, IsTestMode: true}
	service := expense.NewService(db)
	handler := &Handler{expenseService: service}

	httpHandler.Use(middleware.Recover())

	registerRoutes(httpHandler, handler)
}

func registerRoutes(httpHandler *echo.Echo, handler *Handler) {

	httpHandler.POST("/expenses", handler.createHandler)
	httpHandler.GET("/expenses/:id", handler.getHandler)
}
