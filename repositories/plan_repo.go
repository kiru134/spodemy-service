package repositories

import (
	"spodemy-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PlanRepository handles DB operations for Plan.
type PlanRepository struct {
    db *gorm.DB
}

// NewPlanRepository constructs a PlanRepository.
func NewPlanRepository(db *gorm.DB) *PlanRepository {
    return &PlanRepository{db: db}
}

// FindAll returns all plans with their offers preloaded.
func (r *PlanRepository) FindAll() ([]models.Plan, error) {
    var plans []models.Plan
    if err := r.db.Preload("Offers").Find(&plans).Error; err != nil {
        return nil, err
    }
    return plans, nil
}

// FindByID returns one plan by UUID.
func (r *PlanRepository) FindByID(id uuid.UUID) (*models.Plan, error) {
    var plan models.Plan
    if err := r.db.Preload("Offers").First(&plan, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &plan, nil
}

// Create inserts a new plan.
func (r *PlanRepository) Create(p *models.Plan) error {
    return r.db.Create(p).Error
}

// Update modifies an existing plan.
func (r *PlanRepository) Update(p *models.Plan) error {
    return r.db.Save(p).Error
}

// Delete removes a plan by UUID.
func (r *PlanRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.Plan{}, "id = ?", id).Error
}

// AttachOffer associates an offer to a plan.
func (r *PlanRepository) AttachOffer(planID, offerID uuid.UUID) error {
    plan, err := r.FindByID(planID)
    if err != nil {
        return err
    }
    return r.db.Model(plan).Association("Offers").Append(&models.Offer{ID: offerID})
}

// DetachOffer removes the association of an offer from a plan.
func (r *PlanRepository) DetachOffer(planID, offerID uuid.UUID) error {
    plan, err := r.FindByID(planID)
    if err != nil {
        return err
    }
    return r.db.Model(plan).Association("Offers").Delete(&models.Offer{ID: offerID})
}