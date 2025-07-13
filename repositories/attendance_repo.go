package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AttendanceRepository handles DB operations for Attendance.
type AttendanceRepository struct {
    db *gorm.DB
}

// NewAttendanceRepository constructs an AttendanceRepository.
func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
    return &AttendanceRepository{db: db}
}

// FindAll returns all attendance records with enrollment preloaded.
func (r *AttendanceRepository) FindAll() ([]models.Attendance, error) {
    var recs []models.Attendance
    if err := r.db.Preload("Enrollment").Find(&recs).Error; err != nil {
        return nil, err
    }
    return recs, nil
}

// FindByID returns one attendance record by UUID.
func (r *AttendanceRepository) FindByID(id uuid.UUID) (*models.Attendance, error) {
    var a models.Attendance
    if err := r.db.Preload("Enrollment").First(&a, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &a, nil
}

// FindByEnrollment returns records for a specific enrollment UUID.
func (r *AttendanceRepository) FindByEnrollment(enrID uuid.UUID) ([]models.Attendance, error) {
    var recs []models.Attendance
    if err := r.db.Where("enrollment_id = ?", enrID).Preload("Enrollment").Find(&recs).Error; err != nil {
        return nil, err
    }
    return recs, nil
}

// Create inserts a new attendance record.
func (r *AttendanceRepository) Create(a *models.Attendance) error {
    return r.db.Create(a).Error
}

// Update saves changes to an existing attendance record.
func (r *AttendanceRepository) Update(a *models.Attendance) error {
    return r.db.Save(a).Error
}

// Delete removes an attendance record by UUID.
func (r *AttendanceRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.Attendance{}, "id = ?", id).Error
}
