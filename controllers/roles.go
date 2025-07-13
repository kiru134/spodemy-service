package controllers

import (
	"net/http"

	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RoleController handles HTTP requests for roles.
type RoleController struct {
	service *services.RoleService
}

// NewRoleController constructs a RoleController with a service.
func NewRoleController(s *services.RoleService) *RoleController {
	return &RoleController{service: s}
}

// List godoc
// @Summary      List roles
// @Description  Get a list of all roles
// @Tags         roles
// @Produce      json
// @Success      200 {array} models.Role
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /roles [get]
func (ctrl *RoleController) List(c *gin.Context) {
	roles, err := ctrl.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// Get godoc
// @Summary      Get a role
// @Description  Retrieve a role by UUID
// @Tags         roles
// @Produce      json
// @Param        id path string true "Role ID (UUID)"
// @Success      200 {object} models.Role
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /roles/{id} [get]
func (ctrl *RoleController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}
	role, err := ctrl.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

// Create godoc
// @Summary      Create a role
// @Description  Create a new role
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        role body models.Role true "Role object"
// @Success      201 {object} models.Role
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /roles [post]
func (ctrl *RoleController) Create(c *gin.Context) {
	var r models.Role
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.Create(&r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, r)
}

// Update godoc
// @Summary      Update a role
// @Description  Modify an existing role
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        id path string true "Role ID (UUID)"
// @Param        role body models.Role true "Updated role object"
// @Success      200 {object} models.Role
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /roles/{id} [put]
func (ctrl *RoleController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}
	var r models.Role
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r.ID = id
	if err := ctrl.service.Update(&r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

// Delete godoc
// @Summary      Delete a role
// @Description  Remove a role by UUID
// @Tags         roles
// @Param        id path string true "Role ID (UUID)"
// @Success      204 {string} string ""
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     ApiKeyAuth
// @Router       /roles/{id} [delete]
func (ctrl *RoleController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
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