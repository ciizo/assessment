package database

import (
	"testing"

	"github.com/ciizo/assessment/model"
	"github.com/ciizo/assessment/share"
)

func TestCreateExpenseSuccess(t *testing.T) {
	mock := &share.MockDB{}
	db := &Db{DB: mock, IsTestMode: true}
	entity := &model.Expense{Title: "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"}}

	err := db.CreateExpense(entity)

	if err != nil {
		t.Error(err)
	}

}
