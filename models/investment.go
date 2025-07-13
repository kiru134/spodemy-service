package models

import (
	"time"

	"github.com/google/uuid"
)

// Investment represents fractional ownership.
type Investment struct {
    ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    VenueID       uuid.UUID `gorm:"type:uuid;not null;index" json:"venue_id"`
    Venue         Venue     `json:"venue"`
    InvestorID    uuid.UUID `gorm:"type:uuid;not null;index" json:"investor_id"`
    Investor      User      `json:"investor"`
    Units         int       `json:"units"`
    AvgPriceCents int       `json:"avg_price_cents"`
    CreatedAt     time.Time `json:"created_at"`
}

// InvestmentTransaction tracks buy/sell actions.
type InvestmentTransaction struct {
    ID             uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    InvestmentID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"investment_id"`
    Investment     Investment `json:"investment"`
    Type           string     `json:"type"`
    Units          int        `json:"units"`
    PriceCents     int        `json:"price_cents"`
    TransactionRef string     `json:"transaction_ref"`
    TxnDate        time.Time  `json:"txn_date"`
}