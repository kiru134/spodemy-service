package controllers

import (
	"net/http"

	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ExpenseController handles HTTP requests for expenses.
type ExpenseController struct {
    service *services.ExpenseService
}

// NewExpenseController constructs an ExpenseController.
func NewExpenseController(s *services.ExpenseService) *ExpenseController {
    return &ExpenseController{service: s}
}

// List godoc
// @Summary      List expenses
// @Tags         expenses
// @Produce      json
// @Success      200 {array} models.Expense
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /expenses [get]
func (ctrl *ExpenseController) List(c *gin.Context) {
    exps, err := ctrl.service.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, exps)
}

// Get godoc
// @Summary      Get an expense
// @Tags         expenses
// @Produce      json
// @Param        id path string true "Expense UUID"
// @Success      200 {object} models.Expense
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /expenses/{id} [get]
func (ctrl *ExpenseController) Get(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    exp, err := ctrl.service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
        return
    }
    c.JSON(http.StatusOK, exp)
}

// Create godoc
// @Summary      Create an expense
// @Tags         expenses
// @Accept       json
// @Produce      json
// @Param        expense body models.Expense true "Expense object"
// @Success      201 {object} models.Expense
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /expenses [post]
func (ctrl *ExpenseController) Create(c *gin.Context) {
    var e models.Expense
    if err := c.ShouldBindJSON(&e); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.service.Create(&e); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, e)
}

// Update godoc
// @Summary      Update an expense
// @Tags         expenses
// @Accept       json
// @Produce      json
// @Param        id path string true "Expense UUID"
// @Param        expense body models.Expense true "Updated expense object"
// @Success      200 {object} models.Expense
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /expenses/{id} [put]
func (ctrl *ExpenseController) Update(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var e models.Expense
    if err := c.ShouldBindJSON(&e); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    e.ID = id
    if err := ctrl.service.Update(&e); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, e)
}

// Delete godoc
// @Summary      Delete an expense
// @Tags         expenses
// @Param        id path string true "Expense UUID"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /expenses/{id} [delete]
func (ctrl *ExpenseController) Delete(c *gin.Context) {
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