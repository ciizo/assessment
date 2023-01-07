package expense

import (
	"strings"

	"github.com/ciizo/assessment/model"
	"github.com/ciizo/assessment/share"
)

func (service *ExpenseService) Update(id int, entity *model.Expense) error {
	entity.ID = id
	entity.Title = strings.TrimSpace(entity.Title)
	err := share.Validate.Struct(entity)
	if err != nil {
		return model.ServiceErr{Code: share.Error_Invalid_Model, Message: err.Error()}
	}

	err = service.db.UpdateExpense(entity)
	if err != nil {
		return err
	}
	return nil
}
