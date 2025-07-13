package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ExpenseRepository handles DB operations for Expense.
type ExpenseRepository struct {
    db *gorm.DB
}

// NewExpenseRepository constructs an ExpenseRepository.
func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {
    return &ExpenseRepository{db: db}
}

// FindAll returns all expenses.
func (r *ExpenseRepository) FindAll() ([]models.Expense, error) {
    var exps []models.Expense
    if err := r.db.Find(&exps).Error; err != nil {
        return nil, err
    }
    return exps, nil
}

// FindByID returns one expense by UUID.
func (r *ExpenseRepository) FindByID(id uuid.UUID) (*models.Expense, error) {
    var e models.Expense
    if err := r.db.First(&e, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &e, nil
}

// Create inserts a new expense.
func (r *ExpenseRepository) Create(e *models.Expense) error {
    return r.db.Create(e).Error
}

// Update modifies an existing expense.
func (r *ExpenseRepository) Update(e *models.Expense) error {
    return r.db.Save(e).Error
}

// Delete removes an expense by UUID.
func (r *ExpenseRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.Expense{}, "id = ?", id).Error
}