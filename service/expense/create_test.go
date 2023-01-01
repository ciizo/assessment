package expense

import (
	"testing"

	"github.com/ciizo/assessment/database"
	"github.com/ciizo/assessment/model"
	"github.com/ciizo/assessment/share"
	"github.com/go-playground/validator/v10"
)

var service *ExpenseService

func setUp() {
	share.Validate = validator.New()
	mock := &share.MockDB{}
	db := &database.Db{DB: mock, IsTestMode: true}
	service = NewService(db)
}

func TestCreateSuccess(t *testing.T) {
	setUp()
	entity := &model.Expense{
		Title:  "Test T",
		Amount: 79,
		Note:   "Test N",
		Tags:   []string{"T1", "T2"}}

	err := service.Create(entity)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateInvalidModel(t *testing.T) {
	setUp()
	entity := &model.Expense{
		Title:  "Test T",
		Amount: 1,
		Note:   "Test N",
		Tags:   []string{"T1", "T2"}}
	var err error

	t.Run("Title empty", func(t *testing.T) {
		entity.Title = ""
		entity.Amount = 1

		err = service.Create(entity)

		if err == nil {
			t.Fatal("entity data should error invalid Title.")
		}
	})

	t.Run("Title only whitespace", func(t *testing.T) {
		entity.Title = "   "
		entity.Amount = 1

		err := service.Create(entity)

		if err == nil {
			t.Fatal("Entity data should error invalid Title.")
		}
	})

	t.Run("Amount 0", func(t *testing.T) {
		entity.Title = "Test T"
		entity.Amount = 0

		err := service.Create(entity)

		if err == nil {
			t.Fatal("Entity data should error invalid Amount.")
		}
	})

	t.Run("Amount -1", func(t *testing.T) {
		entity.Title = "Test T"
		entity.Amount = -1

		err := service.Create(entity)

		if err == nil {
			t.Fatal("Entity data should error invalid Amount.")
		}
	})

}
