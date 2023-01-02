package expense

import (
	"net/http"

	"github.com/ciizo/assessment/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) getListHandler(c echo.Context) error {
	results, err := h.expenseService.GetList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
