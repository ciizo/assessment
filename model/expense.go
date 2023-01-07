package model

type Expense struct {
	ID     int      `json:"id" validate:"gt=0"`
	Title  string   `json:"title" validate:"required"`
	Amount float64  `json:"amount" validate:"gt=0"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}
