package routes

import (
	"spodemy-backend/controllers"
	"spodemy-backend/repositories"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterPaymentRoutes wires up payment endpoints.
func RegisterPaymentRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    repo := repositories.NewPaymentRepository(db)
    svc := services.NewPaymentService(repo)
    ctrl := controllers.NewPaymentController(svc)

    rg.GET("/payments", ctrl.List)
    rg.POST("/payments", ctrl.Create)
    rg.GET("/payments/:id", ctrl.Get)
    rg.PUT("/payments/:id", ctrl.Update)
    rg.DELETE("/payments/:id", ctrl.Delete)
    rg.GET("/enrollments/:enrollmentId/payments", ctrl.ListByEnrollment)
}
