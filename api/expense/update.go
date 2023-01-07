package expense

import (
	"net/http"

	"github.com/ciizo/assessment/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) updateHandler(c echo.Context) error {
	id, err := validateParamID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Err{Message: err.Error()})
	}

	expense := model.Expense{}
	err = c.Bind(&expense)
	if err != nil {

		return c.JSON(http.StatusBadRequest, model.Err{Message: err.Error()})
	}

	err = h.expenseService.Update(id, &expense)
	if err != nil {
		if model.IsServiceErr(err) {
			return c.JSON(err.(model.ServiceErr).ToHTTPStatus(), model.Err{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, model.Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, expense)
}
