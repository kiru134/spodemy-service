package controllers

import (
	"net/http"

	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AttendanceController handles attendance.
type AttendanceController struct {
    service *services.AttendanceService
}

// NewAttendanceController constructs an AttendanceController.
func NewAttendanceController(s *services.AttendanceService) *AttendanceController {
    return &AttendanceController{service: s}
}

// List godoc
// @Summary      List attendance records
// @Tags         attendance
// @Produce      json
// @Success      200 {array} models.Attendance
// @Failure      500 {object} map[string]string
// @Router       /attendance [get]
func (ctrl *AttendanceController) List(c *gin.Context) {
    recs, err := ctrl.service.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, recs)
}

// ListByEnrollment godoc
// @Summary      List attendance by enrollment
// @Tags         attendance
// @Produce      json
// @Param        enrollmentId path string true "Enrollment UUID"
// @Success      200 {array} models.Attendance
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /enrollments/{enrollmentId}/attendance [get]
func (ctrl *AttendanceController) ListByEnrollment(c *gin.Context) {
    eid, err := uuid.Parse(c.Param("enrollmentId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    recs, err := ctrl.service.ListByEnrollment(eid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, recs)
}


// Get godoc
// @Summary      Get an attendance record by ID
// @Tags         attendance
// @Produce      json
// @Param        id path string true "Attendance UUID"
// @Success      200 {object} models.Attendance
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /attendance/{id} [get]
func (ctrl *AttendanceController) Get(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    a, err := ctrl.service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }
    c.JSON(http.StatusOK, a)
}

// Create godoc
// @Summary      Create a new attendance record
// @Tags         attendance
// @Accept       json
// @Produce      json
// @Param        attendance body models.Attendance true "Attendance object"
// @Success      201 {object} models.Attendance
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /attendance [post]
func (ctrl *AttendanceController) Create(c *gin.Context) {
    var a models.Attendance
    if err := c.ShouldBindJSON(&a); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.service.Create(&a); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, a)
}

// Update godoc
// @Summary      Update an attendance record
// @Tags         attendance
// @Accept       json
// @Produce      json
// @Param        id path string true "Attendance UUID"
// @Param        attendance body models.Attendance true "Updated attendance object"
// @Success      200 {object} models.Attendance
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /attendance/{id} [put]
func (ctrl *AttendanceController) Update(c *gin.Context) {
    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var a models.Attendance
    if err := c.ShouldBindJSON(&a); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    a.ID = id
    if err := ctrl.service.Update(&a); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, a)
}

// Delete godoc
// @Summary      Delete an attendance record
// @Tags         attendance
// @Param        id path string true "Attendance UUID"
// @Success      204 "No Content"
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /attendance/{id} [delete]
func (ctrl *AttendanceController) Delete(c *gin.Context) {
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