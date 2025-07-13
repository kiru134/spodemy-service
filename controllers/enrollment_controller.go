package controllers

import (
	"net/http"

	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// EnrollmentController handles HTTP requests for enrollment endpoints.
type EnrollmentController struct {
    service *services.EnrollmentService
}

// NewEnrollmentController creates a new controller.
func NewEnrollmentController(s *services.EnrollmentService) *EnrollmentController {
    return &EnrollmentController{service: s}
}

// List godoc
// @Summary      List enrollments
// @Tags         enrollments
// @Produce      json
// @Success      200 {array} models.Enrollment
// @Failure      500 {object} map[string]string
// @Router       /enrollments [get]
func (ctrl *EnrollmentController) List(c *gin.Context) {
    ens, err := ctrl.service.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, ens)
}

// Get godoc
// @Summary      Get an enrollment by ID
// @Tags         enrollments
// @Produce      json
// @Param        id path string true "Enrollment UUID"
// @Success      200 {object} models.Enrollment
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /enrollments/{id} [get]
func (ctrl *EnrollmentController) Get(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    e, err := ctrl.service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "enrollment not found"})
        return
    }
    c.JSON(http.StatusOK, e)
}

// Create godoc
// @Summary      Create a new enrollment
// @Tags         enrollments
// @Accept       json
// @Produce      json
// @Param        enrollment body models.Enrollment true "Enrollment object"
// @Success      201 {object} models.Enrollment
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /enrollments [post]
func (ctrl *EnrollmentController) Create(c *gin.Context) {
    var e models.Enrollment
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
// @Summary      Update an existing enrollment
// @Tags         enrollments
// @Accept       json
// @Produce      json
// @Param        id path string true "Enrollment UUID"
// @Param        enrollment body models.Enrollment true "Updated enrollment object"
// @Success      200 {object} models.Enrollment
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /enrollments/{id} [put]
func (ctrl *EnrollmentController) Update(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var e models.Enrollment
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
// @Summary      Delete an enrollment
// @Tags         enrollments
// @Param        id path string true "Enrollment UUID"
// @Success      204 "No Content"
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /enrollments/{id} [delete]
func (ctrl *EnrollmentController) Delete(c *gin.Context) {
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
