package expense

import (
	"database/sql"

	"github.com/ciizo/assessment/database"
	"github.com/ciizo/assessment/share"
	"github.com/go-playground/validator/v10"
)

var service *ExpenseService

type ExpenseService struct {
	db *database.Db
}

func NewService(db *database.Db) *ExpenseService {
	return &ExpenseService{db: db}
}

func setUp() {
	share.Validate = validator.New()
	mock := &share.MockDB{}
	db := &database.Db{DB: mock, IsTestMode: true}
	service = NewService(db)
}

func setUpByDB(mockDb *sql.DB) {
	share.Validate = validator.New()
	db := &database.Db{DB: mockDb, IsTestMode: true}
	service = NewService(db)
}
