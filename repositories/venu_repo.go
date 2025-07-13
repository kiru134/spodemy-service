package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// VenueRepository provides methods to interact with the Venue model.
type VenueRepository struct {
	db *gorm.DB
}

// NewVenueRepository creates a new VenueRepository.
func NewVenueRepository(db *gorm.DB) *VenueRepository {
	return &VenueRepository{db: db}
}

//FindAll returns all venues
func (r*VenueRepository) FindAll() ([]models.Venue,error){
	var venues []models.Venue
	if err := r.db.Find(&venues).Error; err !=nil{
		return nil,err
	}
	return venues,nil
}

// FindByID returns a venue by its ID.
func (r *VenueRepository) FindByID(id uuid.UUID) (*models.Venue, error) {
	var venue models.Venue
	if err := r.db.First(&venue,id).Error; err != nil{
		return nil,err
	}
	return &venue,nil
}

func (r *VenueRepository) Create(venue *models.Venue) error {
    return r.db.Create(venue).Error;
}

// Update modifies an existing venue.
func (r *VenueRepository) Update(v *models.Venue) error {
    return r.db.Save(v).Error
}

// Delete removes a venue by ID.
func (r *VenueRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.Venue{}, id).Error
}
