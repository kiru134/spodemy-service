package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

// RoleService encapsulates business logic for roles.
type RoleService struct {
	repo *repositories.RoleRepository
}

// NewRoleService creates a new RoleService.
func NewRoleService(r *repositories.RoleRepository) *RoleService {
	return &RoleService{repo: r}
}

// List returns all roles.
func (s *RoleService) List() ([]models.Role, error) {
	return s.repo.FindAll()
}

// Get retrieves a single role by UUID.
func (s *RoleService) Get(id uuid.UUID) (*models.Role, error) {
	return s.repo.FindByID(id)
}

// Create adds a new role.
func (s *RoleService) Create(role *models.Role) error {
	return s.repo.Create(role)
}

// Update modifies a role.
func (s *RoleService) Update(role *models.Role) error {
	return s.repo.Update(role)
}

// Delete removes a role.
func (s *RoleService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
