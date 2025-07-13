package models

import (
	"time"

	"github.com/google/uuid"
)

// Role represents a permission or capability.
type Role struct {
    ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    Name      string     `gorm:"unique;not null" json:"name"`
    Users     []*User    `gorm:"many2many:user_roles;" json:"-"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
}

// User represents an application user.
type User struct {
    ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    FirstName    string     `json:"first_name"`
    LastName     string     `json:"last_name"`
    Email        string     `gorm:"unique;not null" json:"email"`
    PasswordHash string     `gorm:"not null" json:"-"`
    Roles        []*Role    `gorm:"many2many:user_roles;" json:"roles"`
    CreatedAt    time.Time  `json:"created_at"`
    UpdatedAt    time.Time  `json:"updated_at"`
}