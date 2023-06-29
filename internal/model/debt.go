package model

type Debt struct {
	ID       int     `json:"id"`
	Amount   float64 `json:"amount"`
	CashBack float64 `json:"cash_back"`
}
