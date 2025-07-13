package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

// AttendanceService encapsulates business logic for attendance.
type AttendanceService struct {
    repo *repositories.AttendanceRepository
}

// NewAttendanceService creates a new AttendanceService.
func NewAttendanceService(r *repositories.AttendanceRepository) *AttendanceService {
    return &AttendanceService{repo: r}
}

// List returns all attendance records.
func (s *AttendanceService) List() ([]models.Attendance, error) {
    return s.repo.FindAll()
}

// ListByEnrollment returns attendance records by enrollment UUID.
func (s *AttendanceService) ListByEnrollment(enrID uuid.UUID) ([]models.Attendance, error) {
    return s.repo.FindByEnrollment(enrID)
}

// Get retrieves a single attendance record by UUID.
func (s *AttendanceService) Get(id uuid.UUID) (*models.Attendance, error) {
    return s.repo.FindByID(id)
}

// Create adds a new attendance record.
func (s *AttendanceService) Create(a *models.Attendance) error {
    return s.repo.Create(a)
}

// Update modifies an attendance record.
func (s *AttendanceService) Update(a *models.Attendance) error {
    return s.repo.Update(a)
}

// Delete removes an attendance record by UUID.
func (s *AttendanceService) Delete(id uuid.UUID) error {
    return s.repo.Delete(id)
}