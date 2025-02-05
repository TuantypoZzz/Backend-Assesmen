package entity

import (
	"github.com/google/uuid"
)

type Deposit struct {
	Id          uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36)"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Username    string    `gorm:"column:username"`
}
