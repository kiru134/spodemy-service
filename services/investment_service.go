package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

// InvestmentService encapsulates logic for investments.
type InvestmentService struct {
    repo *repositories.InvestmentRepository
}

// NewInvestmentService creates a new InvestmentService.
func NewInvestmentService(r *repositories.InvestmentRepository) *InvestmentService {
    return &InvestmentService{repo: r}
}

// List returns all investments.
func (s *InvestmentService) List() ([]models.Investment, error) {
    return s.repo.FindAll()
}

// Get retrieves a single investment.
func (s *InvestmentService) Get(id uuid.UUID) (*models.Investment, error) {
    return s.repo.FindByID(id)
}

// Create adds a new investment.
func (s *InvestmentService) Create(inv *models.Investment) error {
    return s.repo.Create(inv)
}

// Update modifies an investment.
func (s *InvestmentService) Update(inv *models.Investment) error {
    return s.repo.Update(inv)
}

// Delete removes an investment.
func (s *InvestmentService) Delete(id uuid.UUID) error {
    return s.repo.Delete(id)
}

// ListTransactions lists all transactions for an investment.
func (s *InvestmentService) ListTransactions(invID uuid.UUID) ([]models.InvestmentTransaction, error) {
    return s.repo.FindTransactions(invID)
}

// CreateTransaction adds a new transaction.
func (s *InvestmentService) CreateTransaction(txn *models.InvestmentTransaction) error {
    return s.repo.CreateTransaction(txn)
}
