package expense

import (
	"github.com/ciizo/assessment/model"
)

func (service *ExpenseService) GetList() (*[]model.Expense, error) {

	results, err := service.db.GetExpenses()
	if err != nil {
		return nil, err
	}
	return results, nil
}
