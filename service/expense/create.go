package expense

import (
	"strings"

	"github.com/ciizo/assessment/database"
	"github.com/ciizo/assessment/model"
	"github.com/ciizo/assessment/share"
)

type ExpenseService struct {
	db *database.Db
}

func NewService(db *database.Db) *ExpenseService {
	return &ExpenseService{db: db}
}

func (service *ExpenseService) Create(entity *model.Expense) error {
	entity.Title = strings.TrimSpace(entity.Title)
	err := share.Validate.Struct(entity)
	if err != nil {
		return err
	}

	err = service.db.CreateExpense(entity)
	if err != nil {
		return err
	}
	return nil
}
