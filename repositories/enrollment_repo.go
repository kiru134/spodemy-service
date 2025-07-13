package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EnrollmentRepository handles DB operations for Enrollment.
type EnrollmentRepository struct {
    db *gorm.DB
}

// NewEnrollmentRepository constructs an EnrollmentRepository.
func NewEnrollmentRepository(db *gorm.DB) *EnrollmentRepository {
    return &EnrollmentRepository{db: db}
}

// FindAll returns all enrollments with student and batch preloaded.
func (r *EnrollmentRepository) FindAll() ([]models.Enrollment, error) {
    var ens []models.Enrollment
    if err := r.db.Preload("Student").Preload("Batch").Find(&ens).Error; err != nil {
        return nil, err
    }
    return ens, nil
}

// FindByID returns a single enrollment by UUID.
func (r *EnrollmentRepository) FindByID(id uuid.UUID) (*models.Enrollment, error) {
    var e models.Enrollment
    if err := r.db.Preload("Student").Preload("Batch").First(&e, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &e, nil
}

// FindByBatch returns enrollments for a given batch UUID.
func (r *EnrollmentRepository) FindByBatch(batchID uuid.UUID) ([]models.Enrollment, error) {
    var ens []models.Enrollment
    if err := r.db.Where("batch_id = ?", batchID).Preload("Student").Preload("Batch").Find(&ens).Error; err != nil {
        return nil, err
    }
    return ens, nil
}

// Create inserts a new enrollment record.
func (r *EnrollmentRepository) Create(e *models.Enrollment) error {
    return r.db.Create(e).Error
}

// Update saves changes to an existing enrollment.
func (r *EnrollmentRepository) Update(e *models.Enrollment) error {
    return r.db.Save(e).Error
}

// Delete removes an enrollment by UUID.
func (r *EnrollmentRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.Enrollment{}, "id = ?", id).Error
}