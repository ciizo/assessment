package model

import (
	"net/http"

	"github.com/ciizo/assessment/share"
)

func IsServiceErr(t interface{}) bool {
	switch t.(type) {
	case ServiceErr:
		return true
	default:
		return false
	}
}

type ServiceErr struct {
	Code    string
	Message string
}

func (e ServiceErr) Error() string {
	return e.Message
}

func (e ServiceErr) ToHTTPStatus() int {
	switch e.Code {
	case share.Error_Invalid_Model:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
