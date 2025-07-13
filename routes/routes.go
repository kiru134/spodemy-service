package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes registers all API v1 routes on the Gin engine.
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    api := r.Group("/api/v1")

    // venue endpoints
    RegisterVenueRoutes(api, db)

    // user endpoints
    RegisterUserRoutes(api, db)

    RegisterRoleRoutes(api, db)
    RegisterBatchRoutes(api, db)
    RegisterPaymentRoutes(api, db)
    RegisterInvestmentRoutes(api, db)
    RegisterOfferRoutes(api, db)
    RegisterPlanRoutes(api, db)
    RegisterExpenseRoutes(api, db)
    RegisterEnrollmentRoutes(api, db)
    RegisterAttendanceRoutes(api, db)
}
