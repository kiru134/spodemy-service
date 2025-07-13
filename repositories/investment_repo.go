package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// InvestmentRepository handles DB operations for investments & transactions.
type InvestmentRepository struct {
    db *gorm.DB
}

// NewInvestmentRepository constructs an InvestmentRepository.
func NewInvestmentRepository(db *gorm.DB) *InvestmentRepository {
    return &InvestmentRepository{db: db}
}

// FindAll returns all investments.
func (r *InvestmentRepository) FindAll() ([]models.Investment, error) {
    var invs []models.Investment
    if err := r.db.Preload("Venue").Preload("Investor").Find(&invs).Error; err != nil {
        return nil, err
    }
    return invs, nil
}

// FindByID returns one investment by UUID.
func (r *InvestmentRepository) FindByID(id uuid.UUID) (*models.Investment, error) {
    var inv models.Investment
    if err := r.db.Preload("Venue").Preload("Investor").First(&inv, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &inv, nil
}

// Create inserts a new investment.
func (r *InvestmentRepository) Create(inv *models.Investment) error {
    return r.db.Create(inv).Error
}

// Update modifies an existing investment.
func (r *InvestmentRepository) Update(inv *models.Investment) error {
    return r.db.Save(inv).Error
}

// Delete removes an investment by UUID.
func (r *InvestmentRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.Investment{}, "id = ?", id).Error
}

// FindTransactions returns transactions for an investment.
func (r *InvestmentRepository) FindTransactions(invID uuid.UUID) ([]models.InvestmentTransaction, error) {
    var txns []models.InvestmentTransaction
    if err := r.db.Where("investment_id = ?", invID).Find(&txns).Error; err != nil {
        return nil, err
    }
    return txns, nil
}

// CreateTransaction adds a new transaction linked to an investment.
func (r *InvestmentRepository) CreateTransaction(txn *models.InvestmentTransaction) error {
    return r.db.Create(txn).Error
}
