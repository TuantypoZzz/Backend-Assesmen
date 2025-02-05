package model

type CreateTransaction struct {
	AccountNumber string  `json:"account_number"`
	Amount        float64 `json:"amount"`
}

type Transaction struct {
	Saldo float64 `json:"saldo"`
}
