package routes

import (
	"spodemy-backend/controllers"
	"spodemy-backend/repositories"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoleRoutes wires up the /roles endpoints
func RegisterRoleRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewRoleRepository(db)
	svc := services.NewRoleService(repo)
	ctrl := controllers.NewRoleController(svc)

	roles := rg.Group("/roles")
	{
		roles.GET("", ctrl.List)
		roles.GET("/:id", ctrl.Get)
		roles.POST("", ctrl.Create)
		roles.PUT("/:id", ctrl.Update)
		roles.DELETE("/:id", ctrl.Delete)
	}
}