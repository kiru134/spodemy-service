package services

import (
	"spodemy-backend/models"
	"spodemy-backend/repositories"

	"github.com/google/uuid"
)

// PlanService encapsulates business logic for plans.
type PlanService struct {
    repo *repositories.PlanRepository
}

// NewPlanService creates a new PlanService.
func NewPlanService(repo *repositories.PlanRepository) *PlanService {
    return &PlanService{repo: repo}
}

// List returns all plans.
func (s *PlanService) List() ([]models.Plan, error) {
    return s.repo.FindAll()
}

// Get retrieves a single plan.
func (s *PlanService) Get(id uuid.UUID) (*models.Plan, error) {
    return s.repo.FindByID(id)
}

// Create adds a new plan.
func (s *PlanService) Create(plan *models.Plan) error {
    return s.repo.Create(plan)
}

// Update modifies a plan.
func (s *PlanService) Update(plan *models.Plan) error {
    return s.repo.Update(plan)
}

// Delete removes a plan.
func (s *PlanService) Delete(id uuid.UUID) error {
    return s.repo.Delete(id)
}

// AttachOffer associates an offer with a plan.
func (s *PlanService) AttachOffer(planID, offerID uuid.UUID) error {
    return s.repo.AttachOffer(planID, offerID)
}

// DetachOffer removes the association.
func (s *PlanService) DetachOffer(planID, offerID uuid.UUID) error {
    return s.repo.DetachOffer(planID, offerID)
}