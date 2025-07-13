package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OfferRepository handles DB operations for Offer.
type OfferRepository struct {
    db *gorm.DB
}

// NewOfferRepository constructs an OfferRepository.
func NewOfferRepository(db *gorm.DB) *OfferRepository {
    return &OfferRepository{db: db}
}

// FindAll returns all offers.
func (r *OfferRepository) FindAll() ([]models.Offer, error) {
    var offers []models.Offer
    if err := r.db.Preload("Plans").Find(&offers).Error; err != nil {
        return nil, err
    }
    return offers, nil
}

// FindByID returns one offer by UUID.
func (r *OfferRepository) FindByID(id uuid.UUID) (*models.Offer, error) {
    var offer models.Offer
    if err := r.db.Preload("Plans").First(&offer, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &offer, nil
}

// Create inserts a new offer.
func (r *OfferRepository) Create(o *models.Offer) error {
    return r.db.Create(o).Error
}

// Update modifies an existing offer.
func (r *OfferRepository) Update(o *models.Offer) error {
    return r.db.Save(o).Error
}

// Delete removes an offer by UUID.
func (r *OfferRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.Offer{}, "id = ?", id).Error
}