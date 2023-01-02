package expense

import (
	"net/http"
	"strconv"

	"github.com/ciizo/assessment/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) updateHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Err{Message: "id should be int " + err.Error()})
	}
	if id <= 0 {
		return c.JSON(http.StatusBadRequest, model.Err{Message: "id must be greater than 0"})
	}

	expense := model.Expense{}
	err = c.Bind(&expense)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Err{Message: err.Error()})
	}

	err = h.expenseService.Update(id, &expense)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, expense)
}
