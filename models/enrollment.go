package models

import (
	"time"

	"github.com/google/uuid"
)

// Enrollment ties a Student (User) to a Batch.
type Enrollment struct {
    ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    StudentID  uuid.UUID `gorm:"type:uuid;not null;index" json:"student_id"`
    Student    User      `json:"student"`
    BatchID    uuid.UUID `gorm:"type:uuid;not null;index" json:"batch_id"`
    Batch      Batch     `json:"batch"`
    EnrolledOn time.Time `json:"enrolled_on"`
    Status     string    `json:"status"` // "active","completed","dropped"
}

// Attendance per enrollment per date.
type Attendance struct {
    ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    EnrollmentID uuid.UUID  `gorm:"type:uuid;not null;index" json:"enrollment_id"`
    Enrollment   Enrollment `json:"enrollment"`
    Date         time.Time  `json:"date"`
    Status       string     `json:"status"` // "present","absent"
}
