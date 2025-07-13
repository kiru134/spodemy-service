// repositories/role_repo.go
package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleRepository handles DB operations for Role.
type RoleRepository struct {
	db *gorm.DB
}

// NewRoleRepository constructs a RoleRepository with a GORM DB.
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// FindAll returns all roles.
func (r *RoleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// FindByID returns a role by its UUID.
func (r *RoleRepository) FindByID(id uuid.UUID) (*models.Role, error) {
	var role models.Role
	if err := r.db.First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// Create inserts a new role.
func (r *RoleRepository) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

// Update modifies an existing role.
func (r *RoleRepository) Update(role *models.Role) error {
	return r.db.Save(role).Error
}

// Delete removes a role by UUID.
func (r *RoleRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Role{}, "id = ?", id).Error
}