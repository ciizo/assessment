package expense

import (
	"net/http"

	"github.com/ciizo/assessment/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) createHandler(c echo.Context) error {
	expense := model.Expense{}
	err := c.Bind(&expense)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Err{Message: err.Error()})
	}

	err = h.expenseService.Create(&expense)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, expense)
}
