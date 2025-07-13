package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

// EnrollmentService provides business logic for enrollments.
type EnrollmentService struct {
    repo *repositories.EnrollmentRepository
}

// NewEnrollmentService creates a new service instance.
func NewEnrollmentService(r *repositories.EnrollmentRepository) *EnrollmentService {
    return &EnrollmentService{repo: r}
}

// List returns all enrollments.
func (s *EnrollmentService) List() ([]models.Enrollment, error) {
    return s.repo.FindAll()
}

// ListByBatch returns enrollments filtered by batch UUID.
func (s *EnrollmentService) ListByBatch(batchID uuid.UUID) ([]models.Enrollment, error) {
    return s.repo.FindByBatch(batchID)
}

// Get retrieves a single enrollment by UUID.
func (s *EnrollmentService) Get(id uuid.UUID) (*models.Enrollment, error) {
    return s.repo.FindByID(id)
}

// Create adds a new enrollment.
func (s *EnrollmentService) Create(e *models.Enrollment) error {
    return s.repo.Create(e)
}

// Update modifies an existing enrollment.
func (s *EnrollmentService) Update(e *models.Enrollment) error {
    return s.repo.Update(e)
}

// Delete removes an enrollment by UUID.
func (s *EnrollmentService) Delete(id uuid.UUID) error {
    return s.repo.Delete(id)
}
