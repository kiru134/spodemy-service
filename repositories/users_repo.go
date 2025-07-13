package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository handles DB operations for User.
type UserRepository struct {
    db *gorm.DB
}

// NewUserRepository constructs a UserRepository with a GORM DB.
func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

// FindAll returns all users.
func (r *UserRepository) FindAll() ([]models.User, error) {
    var users []models.User
    if err := r.db.Preload("Roles").Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

// FindByID returns a user by its UUID.
func (r *UserRepository) FindByID(id uuid.UUID) (*models.User, error) {
    var user models.User
    if err := r.db.Preload("Roles").First(&user, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// Create inserts a new user record.
func (r *UserRepository) Create(u *models.User) error {
    return r.db.Create(u).Error
}

// Update modifies an existing user.
func (r *UserRepository) Update(u *models.User) error {
    return r.db.Save(u).Error
}

// Delete removes a user by UUID.
func (r *UserRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.User{}, "id = ?", id).Error
}