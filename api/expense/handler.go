package expense

import (
	"github.com/ciizo/assessment/database"

	"github.com/ciizo/assessment/service/expense"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	expenseService *expense.ExpenseService
}

func RegisterHandler(httpHandler *echo.Echo) {

	db := database.NewDb()
	s := expense.NewService(db)
	h := &Handler{expenseService: s}

	httpHandler.Use(middleware.Logger())
	httpHandler.Use(middleware.Recover())

	httpHandler.POST("/expenses", h.createHandler)
}
