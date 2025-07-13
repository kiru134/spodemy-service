package controllers

import (
	"net/http"

	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PaymentController handles HTTP for fee payments.
type PaymentController struct {
    service *services.PaymentService
}

// NewPaymentController constructs a PaymentController.
func NewPaymentController(s *services.PaymentService) *PaymentController {
    return &PaymentController{service: s}
}

// List godoc
// @Summary      List all payments
// @Tags         payments
// @Produce      json
// @Success      200 {array} models.FeePayment
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /payments [get]
func (ctrl *PaymentController) List(c *gin.Context) {
    pmts, err := ctrl.service.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, pmts)
}

// ListByEnrollment godoc
// @Summary      List payments by enrollment
// @Tags         payments
// @Produce      json
// @Param        enrollmentId path string true "Enrollment UUID"
// @Success      200 {array} models.FeePayment
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /enrollments/{enrollmentId}/payments [get]
func (ctrl *PaymentController) ListByEnrollment(c *gin.Context) {
    enrID, err := uuid.Parse(c.Param("enrollmentId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    pmts, err := ctrl.service.ListByEnrollment(enrID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, pmts)
}

// Get godoc
// @Summary      Get a payment
// @Tags         payments
// @Produce      json
// @Param        id path string true "Payment UUID"
// @Success      200 {object} models.FeePayment
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /payments/{id} [get]
func (ctrl *PaymentController) Get(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    pmt, err := ctrl.service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
        return
    }
    c.JSON(http.StatusOK, pmt)
}

// Create godoc
// @Summary      Create a payment
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        payment body models.FeePayment true "Payment object"
// @Success      201 {object} models.FeePayment
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /payments [post]
func (ctrl *PaymentController) Create(c *gin.Context) {
    var p models.FeePayment
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.service.Create(&p); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, p)
}

// Update godoc
// @Summary      Update a payment
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        id path string true "Payment UUID"
// @Param        payment body models.FeePayment true "Updated payment object"
// @Success      200 {object} models.FeePayment
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /payments/{id} [put]
func (ctrl *PaymentController) Update(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var p models.FeePayment
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    p.ID = id
    if err := ctrl.service.Update(&p); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, p)
}

// Delete godoc
// @Summary      Delete a payment
// @Tags         payments
// @Param        id path string true "Payment UUID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /payments/{id} [delete]
func (ctrl *PaymentController) Delete(c *gin.Context) {
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