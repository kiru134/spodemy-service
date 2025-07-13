package models

import (
	"time"

	"github.com/google/uuid"
)

// Plan defines pricing and duration for enrollment.
type Plan struct {
    ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    Name         string    `json:"name"`
    Description  string    `json:"description"`
    PriceCents   int       `json:"price_cents"`
    DurationDays int       `json:"duration_days"`
    Offers       []*Offer  `gorm:"many2many:plan_offers;" json:"offers,omitempty"`
}

// Offer applies a discount and can be attached to multiple plans.
type Offer struct {
    ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    DiscountPct float64   `json:"discount_pct"`
    ValidFrom   time.Time `json:"valid_from"`
    ValidTo     time.Time `json:"valid_to"`
    Plans       []*Plan   `gorm:"many2many:plan_offers;" json:"plans,omitempty"`
}