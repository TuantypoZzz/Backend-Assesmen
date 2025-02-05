package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	AccountID     uuid.UUID `gorm:"primaryKey;column:account_id;type:varchar(36)"`
	AccountNumber string    `gorm:"type:varchar(100);not null;unique" json:"account_number"`
	AccountName   string    `gorm:"type:varchar(100);not null" json:"account_name"`
	IdentityCard  string    `gorm:"type:varchar(16);not null;unique" json:"identity_card"`
	Phone         string    `gorm:"type:varchar(15);not null;unique" json:"phone"`
	Balance       float64   `gorm:"not null;default:0" json:"balance"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// Error implements error.
func (a Account) Error() string {
	panic("unimplemented")
}
