package routes

import (
	"spodemy-backend/controllers"
	"spodemy-backend/repositories"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterExpenseRoutes wires up expense endpoints.
func RegisterExpenseRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    repo := repositories.NewExpenseRepository(db)
    svc := services.NewExpenseService(repo)
    ctrl := controllers.NewExpenseController(svc)

    rg.GET("/expenses", ctrl.List)
    rg.POST("/expenses", ctrl.Create)
    rg.GET("/expenses/:id", ctrl.Get)
    rg.PUT("/expenses/:id", ctrl.Update)
    rg.DELETE("/expenses/:id", ctrl.Delete)
}