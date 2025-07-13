package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

// PaymentService encapsulates logic for fee payments.
type PaymentService struct {
    repo *repositories.PaymentRepository
}

// NewPaymentService creates a new PaymentService.
func NewPaymentService(r *repositories.PaymentRepository) *PaymentService {
    return &PaymentService{repo: r}
}

// List returns all payments.
func (s *PaymentService) List() ([]models.FeePayment, error) {
    return s.repo.FindAll()
}

// ListByEnrollment returns payments by enrollment.
func (s *PaymentService) ListByEnrollment(enrID uuid.UUID) ([]models.FeePayment, error) {
    return s.repo.FindByEnrollment(enrID)
}

// Get retrieves a single payment.
func (s *PaymentService) Get(id uuid.UUID) (*models.FeePayment, error) {
    return s.repo.FindByID(id)
}

// Create adds a new payment.
func (s *PaymentService) Create(p *models.FeePayment) error {
    return s.repo.Create(p)
}

// Update modifies a payment.
func (s *PaymentService) Update(p *models.FeePayment) error {
    return s.repo.Update(p)
}

// Delete removes a payment.
func (s *PaymentService) Delete(id uuid.UUID) error {
    return s.repo.Delete(id)
}