package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BatchRepository handles DB operations for Batch.
type BatchRepository struct {
    db *gorm.DB
}

// NewBatchRepository constructs a BatchRepository.
func NewBatchRepository(db *gorm.DB) *BatchRepository {
    return &BatchRepository{db: db}
}

// FindAll returns all batches.
func (r *BatchRepository) FindAll() ([]models.Batch, error) {
    var batches []models.Batch
    if err := r.db.Preload("Venue").Find(&batches).Error; err != nil {
        return nil, err
    }
    return batches, nil
}

// FindByID returns a batch by its UUID.
func (r *BatchRepository) FindByID(id uuid.UUID) (*models.Batch, error) {
    var batch models.Batch
    if err := r.db.Preload("Venue").First(&batch, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &batch, nil
}

// FindByVenue returns batches for a specific venue.
func (r *BatchRepository) FindByVenue(venueID uuid.UUID) ([]models.Batch, error) {
    var batches []models.Batch
    if err := r.db.Where("venue_id = ?", venueID).Preload("Venue").Find(&batches).Error; err != nil {
        return nil, err
    }
    return batches, nil
}

// Create inserts a new batch.
func (r *BatchRepository) Create(b *models.Batch) error {
    return r.db.Create(b).Error
}

// Update modifies an existing batch.
func (r *BatchRepository) Update(b *models.Batch) error {
    return r.db.Save(b).Error
}

// Delete removes a batch by UUID.
func (r *BatchRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.Batch{}, "id = ?", id).Error
}
