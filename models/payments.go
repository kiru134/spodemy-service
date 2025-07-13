package models

import (
	"time"

	"github.com/google/uuid"
)

// FeePayment records each fee transaction.
type FeePayment struct {
    ID             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    EnrollmentID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"enrollment_id"`
    Enrollment     Enrollment `json:"enrollment"`
    AmountCents    int        `json:"amount_cents"`
    PaidOn         time.Time  `json:"paid_on"`
    Method         string     `json:"method"`
    TransactionRef string     `json:"transaction_ref"`
}
