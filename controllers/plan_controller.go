package controllers

import (
	"net/http"
	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PlanController handles HTTP requests for plans.
type PlanController struct {
    service *services.PlanService
}

// NewPlanController constructs a PlanController.
func NewPlanController(s *services.PlanService) *PlanController {
    return &PlanController{service: s}
}

// List godoc
// @Summary      List plans
// @Tags         plans
// @Produce      json
// @Success      200 {array} models.Plan
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /plans [get]
func (ctrl *PlanController) List(c *gin.Context) {
    plans, err := ctrl.service.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, plans)
}

// Get godoc
// @Summary      Get a plan
// @Tags         plans
// @Produce      json
// @Param        id path string true "Plan ID (UUID)"
// @Success      200 {object} models.Plan
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /plans/{id} [get]
func (ctrl *PlanController) Get(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    plan, err := ctrl.service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "plan not found"})
        return
    }
    c.JSON(http.StatusOK, plan)
}

// Create godoc
// @Summary      Create a plan
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        plan body models.Plan true "Plan object"
// @Success      201 {object} models.Plan
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /plans [post]
func (ctrl *PlanController) Create(c *gin.Context) {
    var p models.Plan
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
// @Summary      Update a plan
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        id path string true "Plan ID (UUID)"
// @Param        plan body models.Plan true "Updated plan object"
// @Success      200 {object} models.Plan
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /plans/{id} [put]
func (ctrl *PlanController) Update(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var p models.Plan
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
// @Summary      Delete a plan
// @Tags         plans
// @Param        id path string true "Plan ID (UUID)"
// @Success      204 {string} string ""
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /plans/{id} [delete]
func (ctrl *PlanController) Delete(c *gin.Context) {
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