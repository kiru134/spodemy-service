package controllers

import (
	"net/http"
	"spodemy-backend/models"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserController handles HTTP requests for users.
type UserController struct {
    service *services.UserService
}

// NewUserController constructs a UserController with a service.
func NewUserController(s *services.UserService) *UserController {
    return &UserController{service: s}
}

// List godoc
// @Summary      List users
// @Description  Get a list of all users
// @Tags         users
// @Produce      json
// @Success      200  {array}   models.User
// @Failure      500  {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /users [get]
func (ctrl *UserController) List(c *gin.Context) {
    users, err := ctrl.service.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

// Get godoc
// @Summary      Get a user
// @Description  Retrieve a user by UUID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID (UUID)"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /users/{id} [get]
func (ctrl *UserController) Get(c *gin.Context) {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    user, err := ctrl.service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// Create godoc
// @Summary      Create a user
// @Description  Create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User object"
// @Success      201   {object}  models.User
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /users [post]
func (ctrl *UserController) Create(c *gin.Context) {
    var u models.User
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.service.Create(&u); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, u)
}

// Update godoc
// @Summary      Update a user
// @Description  Modify an existing user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      string        true  "User ID (UUID)"
// @Param        user  body      models.User   true  "Updated user object"
// @Success      200   {object}  models.User
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /users/{id} [put]
func (ctrl *UserController) Update(c *gin.Context) {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }
    var u models.User
    if err := c.ShouldBindJSON(&u); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    u.ID = id
    if err := ctrl.service.Update(&u); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, u)
}

// Delete godoc
// @Summary      Delete a user
// @Description  Remove a user by UUID
// @Tags         users
// @Param        id   path      string  true  "User ID (UUID)"
// @Success      204  {string}  string  ""
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /users/{id} [delete]
func (ctrl *UserController) Delete(c *gin.Context) {
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

// routes/routes.go (add User routes)
// Inside SetupRoutes after venue routes:
//
// userRepo := repositories.NewUserRepository(db)
// userSvc := services.NewUserService(userRepo)
// userCtrl := controllers.NewUserController(userSvc)
//
// users := api.Group("/users")
// {
//     users.GET("", userCtrl.List)
//     users.GET(":id", userCtrl.Get)
//     users.POST("", middlewares.Authorize("admin"), userCtrl.Create)
//     users.PUT(":id", middlewares.Authorize("admin"), userCtrl.Update)
//     users.DELETE(":id", middlewares.Authorize("admin"), userCtrl.Delete)
// }
