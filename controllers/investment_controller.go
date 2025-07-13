package controllers

import (
	"net/http"

	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// InvestmentController handles HTTP requests for investments.
type InvestmentController struct {
    service *services.InvestmentService
}

// NewInvestmentController constructs an InvestmentController.
func NewInvestmentController(s *services.InvestmentService) *InvestmentController {
    return &InvestmentController{service: s}
}

// List godoc
// @Summary      List investments
// @Tags         investments
// @Produce      json
// @Success      200 {array} models.Investment
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /investments [get]
func (ctrl *InvestmentController) List(c *gin.Context) {
    invs, err := ctrl.service.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, invs)
}

// Get godoc
// @Summary      Get an investment
// @Tags         investments
// @Produce      json
// @Param        id path string true "Investment UUID"
// @Success      200 {object} models.Investment
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /investments/{id} [get]
func (ctrl *InvestmentController) Get(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    inv, err := ctrl.service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "investment not found"})
        return
    }
    c.JSON(http.StatusOK, inv)
}

// Create godoc
// @Summary      Create an investment
// @Tags         investments
// @Accept       json
// @Produce      json
// @Param        investment body models.Investment true "Investment object"
// @Success      201 {object} models.Investment
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /investments [post]
func (ctrl *InvestmentController) Create(c *gin.Context) {
    var inv models.Investment
    if err := c.ShouldBindJSON(&inv); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.service.Create(&inv); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, inv)
}

// Update godoc
// @Summary      Update an investment
// @Tags         investments
// @Accept       json
// @Produce      json
// @Param        id path string true "Investment UUID"
// @Param        investment body models.Investment true "Updated investment object"
// @Success      200 {object} models.Investment
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /investments/{id} [put]
func (ctrl *InvestmentController) Update(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var inv models.Investment
    if err := c.ShouldBindJSON(&inv); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    inv.ID = id
    if err := ctrl.service.Update(&inv); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, inv)
}

// Delete godoc
// @Summary      Delete an investment
// @Tags         investments
// @Param        id path string true "Investment UUID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /investments/{id} [delete]
func (ctrl *InvestmentController) Delete(c *gin.Context) {
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

// ListTransactions godoc
// @Summary      List transactions for an investment
// @Tags         investments
// @Produce      json
// @Param        id path string true "Investment UUID"
// @Success      200 {array} models.InvestmentTransaction
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /investments/{id}/transactions [get]
func (ctrl *InvestmentController) ListTransactions(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    txns, err := ctrl.service.ListTransactions(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, txns)
}

// CreateTransaction godoc
// @Summary      Create a transaction for an investment
// @Tags         investments
// @Accept       json
// @Produce      json
// @Param        id path string true "Investment UUID"
// @Param        txn body models.InvestmentTransaction true "Transaction object"
// @Success      201 {object} models.InvestmentTransaction
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /investments/{id}/transactions [post]
func (ctrl *InvestmentController) CreateTransaction(c *gin.Context) {
    invID, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var txn models.InvestmentTransaction
    if err := c.ShouldBindJSON(&txn); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    txn.InvestmentID = invID
    if err := ctrl.service.CreateTransaction(&txn); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, txn)
}
