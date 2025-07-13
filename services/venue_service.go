package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

//VenueService encapsulates business logic for venues.

type VenueService struct {
	repo *repositories.VenueRepository
}

// NewVenueService creates a new VenueService.
func NewVenueService(r *repositories.VenueRepository) *VenueService {
    return &VenueService{repo: r}
}

// List returns all venues.
func (s *VenueService) ListAllVenue() ([]models.Venue, error) {
    return s.repo.FindAll()
}

// Get retrieves a single venue by ID.
func (s *VenueService) GetAVenue(id uuid.UUID) (*models.Venue, error) {
    return s.repo.FindByID(id)
}

// Create adds a new venue.
func (s *VenueService) CreateVenue(v *models.Venue) error {
    return s.repo.Create(v)
}

// Update modifies a venue.
func (s *VenueService) UpdateVenue(v *models.Venue) error {
    return s.repo.Update(v)
}

// Delete removes a venue.
func (s *VenueService) DeleteVenue(id uuid.UUID) error {
    return s.repo.Delete(id)
}
