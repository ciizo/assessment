package expense

import (
	"errors"
	"strconv"
)

func validateParamID(textID string) (int, error) {
	id, err := strconv.Atoi(textID)
	if err != nil {
		return 0, errors.New("id should be int " + err.Error())
	}
	if id <= 0 {
		return 0, errors.New("id must be greater than 0")
	}

	return id, nil
}
