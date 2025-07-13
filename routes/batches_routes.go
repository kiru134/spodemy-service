package routes

import (
	"spodemy-backend/controllers"
	"spodemy-backend/repositories"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterBatchRoutes wires up batch endpoints.
func RegisterBatchRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    repo := repositories.NewBatchRepository(db)
    svc := services.NewBatchService(repo)
    ctrl := controllers.NewBatchController(svc)

    // Flat batch routes
    rg.GET("/batches", ctrl.List)
    rg.GET("/batches/:id", ctrl.Get)
    rg.PUT("/batches/:id", ctrl.Update)
    rg.DELETE("/batches/:id", ctrl.Delete)

    // Nested under venues
    rg.GET("/venues/:venueId/batches", ctrl.ListByVenue)
    rg.POST("/venues/:venueId/batches", ctrl.Create)
}