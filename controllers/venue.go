package controllers

import (
	"net/http"
	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// VenueController handles HTTP requests for venues.
type VenueController struct {
    service *services.VenueService
}

// NewVenueController constructs a VenueController with a service.
func NewVenueController(s *services.VenueService) *VenueController {
    return &VenueController{service: s}
}

// List godoc
// @Summary      List venues
// @Description  Get a list of all venues
// @Tags         venues
// @Produce      json
// @Success      200  {array}   models.Venue
// @Failure      500  {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /venues [get]
func (ctrl *VenueController) List(c *gin.Context) {
    venues, err := ctrl.service.ListAllVenue()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, venues)
}

// Get godoc
// @Summary      Get a venue
// @Description  Retrieve a venue by its UUID
// @Tags         venues
// @Produce      json
// @Param        id   path      string  true  "Venue ID (UUID format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)"
// @Success      200  {object}  models.Venue
// @Failure      400  {object}  map[string]string "Invalid UUID format"
// @Failure      404  {object}  map[string]string "Venue not found"
// @Security     ApiKeyAuth
// @Router       /venues/{id} [get]
func (ctrl *VenueController) Get(c *gin.Context) {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    venue, err := ctrl.service.GetAVenue(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "venue not found"})
        return
    }
    c.JSON(http.StatusOK, venue)
}

// Create godoc
// @Summary      Create a venue
// @Description  Create a new venue record
// @Tags         venues
// @Accept       json
// @Produce      json
// @Param        venue  body      models.Venue  true  "Venue object"
// @Success      201    {object}  models.Venue
// @Failure      400    {object}  map[string]string "Invalid request body"
// @Failure      500    {object}  map[string]string "Server error"
// @Security     ApiKeyAuth
// @Router       /venues [post]
func (ctrl *VenueController) Create(c *gin.Context) {
    var v models.Venue
    if err := c.ShouldBindJSON(&v); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.service.CreateVenue(&v); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, v)
}

// Update godoc
// @Summary      Update a venue
// @Description  Modify an existing venue
// @Tags         venues
// @Accept       json
// @Produce      json
// @Param        id     path      string        true  "Venue ID (UUID format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)"
// @Param        venue  body      models.Venue  true  "Updated venue object"
// @Success      200    {object}  models.Venue
// @Failure      400    {object}  map[string]string "Invalid UUID format or request body"
// @Failure      404    {object}  map[string]string "Venue not found"
// @Failure      500    {object}  map[string]string "Server error"
// @Security     ApiKeyAuth
// @Router       /venues/{id} [put]
func (ctrl *VenueController) Update(c *gin.Context) {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var v models.Venue
    if err := c.ShouldBindJSON(&v); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    v.ID = id
    if err := ctrl.service.UpdateVenue(&v); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, v)
}

// Delete godoc
// @Summary      Delete a venue
// @Description  Remove a venue by UUID
// @Tags         venues
// @Param        id   path      string  true  "Venue ID (UUID format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)"
// @Success      204  {string}  string  "No Content"
// @Failure      400  {object}  map[string]string "Invalid UUID format"
// @Failure      404  {object}  map[string]string "Venue not found"
// @Failure      500  {object}  map[string]string "Server error"
// @Security     ApiKeyAuth
// @Router       /venues/{id} [delete]
func (ctrl *VenueController) Delete(c *gin.Context) {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    if err := ctrl.service.DeleteVenue(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
