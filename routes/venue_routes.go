package routes

import (
	"spodemy-backend/controllers"
	"spodemy-backend/repositories"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterVenueRoutes wires up the /venues endpoints under the given router group.
func RegisterVenueRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    repo := repositories.NewVenueRepository(db)
    svc  := services.NewVenueService(repo)
    ctrl := controllers.NewVenueController(svc)

    venues := rg.Group("/venues")
    {
        venues.GET("", ctrl.List)
        venues.GET("/:id", ctrl.Get)
        venues.POST("", ctrl.Create)
        venues.PUT("/:id", ctrl.Update)
        venues.DELETE("/:id", ctrl.Delete)
    }
}