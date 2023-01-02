package expense

import (
	"errors"

	"github.com/ciizo/assessment/model"
)

func (service *ExpenseService) Get(id int) (*model.Expense, error) {
	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}

	result, err := service.db.GetExpense(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
