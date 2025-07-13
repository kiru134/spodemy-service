package routes

import (
	"spodemy-backend/controllers"
	"spodemy-backend/repositories"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterInvestmentRoutes wires up investment endpoints.
func RegisterInvestmentRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    repo := repositories.NewInvestmentRepository(db)
    svc := services.NewInvestmentService(repo)
    ctrl := controllers.NewInvestmentController(svc)

    rg.GET("/investments", ctrl.List)
    rg.POST("/investments", ctrl.Create)
    rg.GET("/investments/:id", ctrl.Get)
    rg.PUT("/investments/:id", ctrl.Update)
    rg.DELETE("/investments/:id", ctrl.Delete)

    rg.GET("/investments/:id/transactions", ctrl.ListTransactions)
    rg.POST("/investments/:id/transactions", ctrl.CreateTransaction)
}
