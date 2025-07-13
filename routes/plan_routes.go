package routes

import (
	"spodemy-backend/controllers"
	"spodemy-backend/repositories"
	"spodemy-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterPlanRoutes sets up plan-related routes
func RegisterPlanRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    repo := repositories.NewPlanRepository(db)
    svc := services.NewPlanService(repo)
    ctrl := controllers.NewPlanController(svc)

    plans := rg.Group("/plans")
    {
        plans.GET("", ctrl.List)
        plans.POST("", ctrl.Create)
        plans.GET("/:id", ctrl.Get)
        plans.PUT("/:id", ctrl.Update)
        plans.DELETE("/:id", ctrl.Delete)
    }
}

// RegisterOfferRoutes sets up offer-related routes
func RegisterOfferRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    repo := repositories.NewOfferRepository(db)
    svc := services.NewOfferService(repo)
    ctrl := controllers.NewOfferController(svc)

    offers := rg.Group("/offers")
    {
        offers.GET("", ctrl.List)
        offers.POST("", ctrl.Create)
        offers.GET("/:id", ctrl.Get)
        offers.PUT("/:id", ctrl.Update)
        offers.DELETE("/:id", ctrl.Delete)
    }
}