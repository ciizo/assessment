package expense

import (
	"net/http"

	"github.com/ciizo/assessment/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) getHandler(c echo.Context) error {
	id, err := validateParamID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Err{Message: err.Error()})
	}

	result := &model.Expense{}
	result, err = h.expenseService.Get(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
