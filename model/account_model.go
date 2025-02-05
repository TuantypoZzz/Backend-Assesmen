package model

type CreateAccount struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name" validate:"required,max=32"`
	IdentityCard  string `json:"identity_card"`
	Phone         string `json:"phone"`
	CreatedAt     string `json:"created_at"`
}

type RegisterAccount struct {
	AccountNumber string `json:"account_number"`
}

type BalanceAccount struct {
	Saldo float64 `json:"balance"`
}
