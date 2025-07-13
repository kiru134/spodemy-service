package routes

import (
	"spodemy-backend/controllers"
	"spodemy-backend/repositories"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUserRoutes wires up the /users endpoints under the given router group.
func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    repo := repositories.NewUserRepository(db)
    svc  := services.NewUserService(repo)
    ctrl := controllers.NewUserController(svc)

    users := rg.Group("/users")
    {
        users.GET("", ctrl.List)
        users.GET("/:id", ctrl.Get)
        users.POST("", ctrl.Create)
        users.PUT("/:id", ctrl.Update)
        users.DELETE("/:id", ctrl.Delete)
    }
}