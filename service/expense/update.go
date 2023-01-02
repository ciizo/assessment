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
		return err
	}

	err = service.db.UpdateExpense(entity)
	if err != nil {
		return err
	}
	return nil
}
