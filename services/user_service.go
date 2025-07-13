package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

// UserService encapsulates business logic for users.
type UserService struct {
    repo *repositories.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(r *repositories.UserRepository) *UserService {
    return &UserService{repo: r}
}

// List returns all users.
func (s *UserService) List() ([]models.User, error) {
    return s.repo.FindAll()
}

// Get retrieves a single user by UUID.
func (s *UserService) Get(id uuid.UUID) (*models.User, error) {
    return s.repo.FindByID(id)
}

// Create adds a new user.
func (s *UserService) Create(u *models.User) error {
    return s.repo.Create(u)
}

// Update modifies a user.
func (s *UserService) Update(u *models.User) error {
    return s.repo.Update(u)
}

// Delete removes a user.
func (s *UserService) Delete(id uuid.UUID) error {
    return s.repo.Delete(id)
}

