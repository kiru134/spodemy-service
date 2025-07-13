package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PaymentRepository handles DB operations for fee payments.
type PaymentRepository struct {
    db *gorm.DB
}

// NewPaymentRepository constructs a PaymentRepository.
func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
    return &PaymentRepository{db: db}
}

// FindAll returns all payments.
func (r *PaymentRepository) FindAll() ([]models.FeePayment, error) {
    var payments []models.FeePayment
    if err := r.db.Preload("Enrollment").Find(&payments).Error; err != nil {
        return nil, err
    }
    return payments, nil
}

// FindByID returns one payment by UUID.
func (r *PaymentRepository) FindByID(id uuid.UUID) (*models.FeePayment, error) {
    var p models.FeePayment
    if err := r.db.Preload("Enrollment").First(&p, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &p, nil
}

// FindByEnrollment returns payments for a specific enrollment.
func (r *PaymentRepository) FindByEnrollment(enrID uuid.UUID) ([]models.FeePayment, error) {
    var payments []models.FeePayment
    if err := r.db.Where("enrollment_id = ?", enrID).Preload("Enrollment").Find(&payments).Error; err != nil {
        return nil, err
    }
    return payments, nil
}

// Create inserts a new payment.
func (r *PaymentRepository) Create(p *models.FeePayment) error {
    return r.db.Create(p).Error
}

// Update modifies an existing payment.
func (r *PaymentRepository) Update(p *models.FeePayment) error {
    return r.db.Save(p).Error
}

// Delete removes a payment by UUID.
func (r *PaymentRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.FeePayment{}, "id = ?", id).Error
}
