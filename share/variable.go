package share

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

const IT_Test_ServerPort = 80
const IT_Test_DB_ConnectionString = "postgresql://root:root@db/go-it-db?sslmode=disable"
const Error_Invalid_Model = "01"
