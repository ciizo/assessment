package expense

import (
	"strings"

	"github.com/ciizo/assessment/model"
	"github.com/ciizo/assessment/share"
)

func (service *ExpenseService) Create(entity *model.Expense) error {
	entity.Title = strings.TrimSpace(entity.Title)
	err := share.Validate.Struct(entity)
	if err != nil {
		return model.ServiceErr{Code: share.Error_Invalid_Model, Message: err.Error()}
	}

	err = service.db.CreateExpense(entity)
	if err != nil {
		return err
	}
	return nil
}
