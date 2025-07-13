package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

// BatchService encapsulates business logic for batches.
type BatchService struct {
    repo *repositories.BatchRepository
}

// NewBatchService creates a new BatchService.
func NewBatchService(r *repositories.BatchRepository) *BatchService {
    return &BatchService{repo: r}
}

// List returns all batches.
func (s *BatchService) List() ([]models.Batch, error) {
    return s.repo.FindAll()
}

// ListByVenue returns batches for a venue.
func (s *BatchService) ListByVenue(venueID uuid.UUID) ([]models.Batch, error) {
    return s.repo.FindByVenue(venueID)
}

// Get retrieves a single batch by UUID.
func (s *BatchService) Get(id uuid.UUID) (*models.Batch, error) {
    return s.repo.FindByID(id)
}

// Create adds a new batch.
func (s *BatchService) Create(b *models.Batch) error {
    return s.repo.Create(b)
}

// Update modifies a batch.
func (s *BatchService) Update(b *models.Batch) error {
    return s.repo.Update(b)
}

// Delete removes a batch.
func (s *BatchService) Delete(id uuid.UUID) error {
    return s.repo.Delete(id)
}