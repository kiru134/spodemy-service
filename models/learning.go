package models

import (
	"time"

	"github.com/google/uuid"
)

// Course created by a Coach.
type Course struct {
    ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    CoachID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"coach_id"`
    Coach       User       `json:"coach"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    ContentURL  string     `json:"content_url"`
    CreatedAt   time.Time  `json:"created_at"`
}

// Assessment records student performance.
type Assessment struct {
    ID          uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
    CourseID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"course_id"`
    Course      Course     `json:"course"`
    StudentID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"student_id"`
    Student     User       `json:"student"`
    Score       float64    `json:"score"`
    AttemptedAt time.Time  `json:"attempted_at"`
}

// // Certification issued upon course completion.
// type Certification struct {
//     ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create" json:"id" binding:"-"`
//     CourseID  uuid.UUID `gorm:"type:uuid;not null;index" json:"course_id"`
//     Course    Course    `json:"course"`
//     StudentID uuid.UUID `gorm:"type:uuid;not null;index" json:"student_id"`
//     Student   User      `json:"student"`
//     IssuedOn  time.Time `json:"issued_on"`
// }