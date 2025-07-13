package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

// OfferService encapsulates business logic for offers.
type OfferService struct {
    repo *repositories.OfferRepository
}

// NewOfferService creates a new OfferService.
func NewOfferService(repo *repositories.OfferRepository) *OfferService {
    return &OfferService{repo: repo}
}

// List returns all offers.
func (s *OfferService) List() ([]models.Offer, error) {
    return s.repo.FindAll()
}

// Get retrieves a single offer.
func (s *OfferService) Get(id uuid.UUID) (*models.Offer, error) {
    return s.repo.FindByID(id)
}

// Create adds a new offer.
func (s *OfferService) Create(offer *models.Offer) error {
    return s.repo.Create(offer)
}

// Update modifies an offer.
func (s *OfferService) Update(offer *models.Offer) error {
    return s.repo.Update(offer)
}

// Delete removes an offer.
func (s *OfferService) Delete(id uuid.UUID) error {
    return s.repo.Delete(id)
}