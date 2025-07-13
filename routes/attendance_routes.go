package routes

import (
	"spodemy-backend/controllers"
	"spodemy-backend/repositories"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterAttendanceRoutes sets up attendance endpoints.
func RegisterAttendanceRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    repo := repositories.NewAttendanceRepository(db)
    svc := services.NewAttendanceService(repo)
    ctrl := controllers.NewAttendanceController(svc)

    att := rg.Group("/attendance")
    {
        att.GET("", ctrl.List)
        att.GET("/:id", ctrl.Get)
        att.POST("", ctrl.Create)
        att.PUT("/:id", ctrl.Update)
        att.DELETE("/:id", ctrl.Delete)
    }

    // nested under enrollments
    rg.GET("/enrollments/:enrollmentId/attendance", ctrl.ListByEnrollment)
}