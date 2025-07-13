package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

// ExpenseService encapsulates business logic for expenses.
type ExpenseService struct {
    repo *repositories.ExpenseRepository
}

// NewExpenseService creates a new ExpenseService.
func NewExpenseService(r *repositories.ExpenseRepository) *ExpenseService {
    return &ExpenseService{repo: r}
}

// List returns all expenses.
func (s *ExpenseService) List() ([]models.Expense, error) {
    return s.repo.FindAll()
}

// Get retrieves a single expense.
func (s *ExpenseService) Get(id uuid.UUID) (*models.Expense, error) {
    return s.repo.FindByID(id)
}

// Create adds a new expense.
func (s *ExpenseService) Create(e *models.Expense) error {
    return s.repo.Create(e)
}

// Update modifies an expense.
func (s *ExpenseService) Update(e *models.Expense) error {
    return s.repo.Update(e)
}

// Delete removes an expense.
func (s *ExpenseService) Delete(id uuid.UUID) error {
    return s.repo.Delete(id)
}
