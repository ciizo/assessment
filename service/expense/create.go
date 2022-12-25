package expense

import (
	"github.com/ciizo/assessment/database"
	"github.com/ciizo/assessment/model"
)

type ExpenseService struct {
	db *database.Db
}

func NewService(db *database.Db) *ExpenseService {
	return &ExpenseService{db: db}
}

func (service *ExpenseService) Create(entity *model.Expense) error {
	err := service.db.CreateExpense(entity)
	if err != nil {
		return err
	}
	return nil
}
