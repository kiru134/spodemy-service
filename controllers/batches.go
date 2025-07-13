package controllers

import (
	"net/http"
	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// BatchController handles HTTP requests for batches.
type BatchController struct {
    service *services.BatchService
}

// NewBatchController constructs a BatchController.
func NewBatchController(s *services.BatchService) *BatchController {
    return &BatchController{service: s}
}

// List godoc
// @Summary      List all batches
// @Tags         batches
// @Produce      json
// @Success      200 {array} models.Batch
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /batches [get]
func (ctrl *BatchController) List(c *gin.Context) {
    batches, err := ctrl.service.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, batches)
}

// ListByVenue godoc
// @Summary      List batches by venue
// @Tags         batches
// @Produce      json
// @Param        venueId  path  string  true  "Venue ID (UUID)"
// @Success      200 {array} models.Batch
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /venues/{venueId}/batches [get]
func (ctrl *BatchController) ListByVenue(c *gin.Context) {
    venueID, err := uuid.Parse(c.Param("venueId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid venue UUID"})
        return
    }
    batches, err := ctrl.service.ListByVenue(venueID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, batches)
}

// Get godoc
// @Summary      Get a batch
// @Tags         batches
// @Produce      json
// @Param        id  path  string  true  "Batch ID (UUID)"
// @Success      200 {object} models.Batch
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /batches/{id} [get]
func (ctrl *BatchController) Get(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    batch, err := ctrl.service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "batch not found"})
        return
    }
    c.JSON(http.StatusOK, batch)
}

// Create godoc
// @Summary      Create a batch for a venue
// @Tags         batches
// @Accept       json
// @Produce      json
// @Param        venueId  path  string      true  "Venue ID (UUID)"
// @Param        batch    body  models.Batch  true  "Batch object"
// @Success      201 {object} models.Batch
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /venues/{venueId}/batches [post]
func (ctrl *BatchController) Create(c *gin.Context) {
    venueID, err := uuid.Parse(c.Param("venueId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid venue UUID"})
        return
    }
    var b models.Batch
    if err := c.ShouldBindJSON(&b); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    b.VenueID = venueID
    if err := ctrl.service.Create(&b); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, b)
}

// Update godoc
// @Summary      Update a batch
// @Tags         batches
// @Accept       json
// @Produce      json
// @Param        id    path  string      true  "Batch ID (UUID)"
// @Param        batch body  models.Batch  true  "Updated batch object"
// @Success      200 {object} models.Batch
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /batches/{id} [put]
func (ctrl *BatchController) Update(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var b models.Batch
    if err := c.ShouldBindJSON(&b); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    b.ID = id
    if err := ctrl.service.Update(&b); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, b)
}

// Delete godoc
// @Summary      Delete a batch
// @Tags         batches
// @Param        id  path  string  true  "Batch ID (UUID)"
// @Success      204 {string} string ""
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /batches/{id} [delete]
func (ctrl *BatchController) Delete(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    if err := ctrl.service.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
