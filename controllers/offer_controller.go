package controllers

import (
	"net/http"
	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// OfferController handles HTTP requests for offers.
type OfferController struct {
    service *services.OfferService
}

// NewOfferController constructs an OfferController.
func NewOfferController(s *services.OfferService) *OfferController {
    return &OfferController{service: s}
}

// List godoc
// @Summary      List offers
// @Tags         offers
// @Produce      json
// @Success      200 {array} models.Offer
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /offers [get]
func (ctrl *OfferController) List(c *gin.Context) {
    offers, err := ctrl.service.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, offers)
}

// Get godoc
// @Summary      Get an offer
// @Tags         offers
// @Produce      json
// @Param        id path string true "Offer ID (UUID)"
// @Success      200 {object} models.Offer
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /offers/{id} [get]
func (ctrl *OfferController) Get(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    offer, err := ctrl.service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "offer not found"})
        return
    }
    c.JSON(http.StatusOK, offer)
}

// Create godoc
// @Summary      Create an offer
// @Tags         offers
// @Accept       json
// @Produce      json
// @Param        offer body models.Offer true "Offer object"
// @Success      201 {object} models.Offer
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /offers [post]
func (ctrl *OfferController) Create(c *gin.Context) {
    var offer models.Offer
    if err := c.ShouldBindJSON(&offer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.service.Create(&offer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, offer)
}

// Update godoc
// @Summary      Update an offer
// @Tags         offers
// @Accept       json
// @Produce      json
// @Param        id path string true "Offer ID (UUID)"
// @Param        offer body models.Offer true "Updated offer object"
// @Success      200 {object} models.Offer
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /offers/{id} [put]
func (ctrl *OfferController) Update(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var offer models.Offer
    if err := c.ShouldBindJSON(&offer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    offer.ID = id
    if err := ctrl.service.Update(&offer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, offer)
}

// Delete godoc
// @Summary      Delete an offer
// @Tags         offers
// @Param        id path string true "Offer ID (UUID)"
// @Success      204 {string} string ""
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /offers/{id} [delete]
func (ctrl *OfferController) Delete(c *gin.Context) {
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