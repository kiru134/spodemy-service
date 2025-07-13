package routes

import (
	"spodemy-backend/controllers"
	"spodemy-backend/repositories"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterEnrollmentRoutes sets up enrollment endpoints.
func RegisterEnrollmentRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    repo := repositories.NewEnrollmentRepository(db)
    svc := services.NewEnrollmentService(repo)
    ctrl := controllers.NewEnrollmentController(svc)

    ens := rg.Group("/enrollments")
    {
        ens.GET("", ctrl.List)
        ens.GET("/:id", ctrl.Get)
        ens.POST("", ctrl.Create)
        ens.PUT("/:id", ctrl.Update)
        ens.DELETE("/:id", ctrl.Delete)
    }

    // nested under batches
    rg.GET("/batches/:batchId/enrollments", ctrl.ListByBatch)
}