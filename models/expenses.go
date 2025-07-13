package models

import (
	"time"

	"github.com/google/uuid"
)

// Expense logged for operational costs.
type Expense struct {
    ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    Description string    `json:"description"`
    AmountCents int       `json:"amount_cents"`
    IncurredOn  time.Time `json:"incurred_on"`
}