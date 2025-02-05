package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID uuid.UUID `gorm:"primaryKey;column:transaction_id;type:varchar(36)"`
	AccountNumber string    `gorm:"foreignKey:account_number"`
	Type          string    `gorm:"type:varchar(10);not null;check:type IN ('deposit', 'withdraw')" json:"type"` // "deposit" or "withdraw"
	Amount        float64   `gorm:"type:decimal(15,2);not null" json:"amount"`
	FinalBalance  float64   `gorm:"type:decimal(15,2);not null" json:"final_balance"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}
