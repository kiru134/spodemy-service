package models

import (
	"time"

	"github.com/google/uuid"
)

// Venue where batches run and investments attach.
type Venue struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    Name      string    `gorm:"not null" json:"name"`
    Location  string    `json:"location"`
    Capacity  int       `json:"capacity"`
    Batches   []Batch   `json:"batches,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// Batch groups students at a Venue.
type Batch struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    VenueID   uuid.UUID `gorm:"type:uuid;not null;index" json:"venue_id"`
    Venue     Venue     `json:"venue"`
    Name      string    `json:"name"`
    StartDate time.Time `json:"start_date"`
    EndDate   time.Time `json:"end_date"`
}
