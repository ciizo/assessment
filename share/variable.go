package share

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

const ITTestDbConnectioString = "postgresql://root:root@db/go-it-db?sslmode=disable"
