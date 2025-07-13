// @BasePath /api/v1
// @title        Spodemy API
// @version      1.0
// @description  API documentation for the Spodemy multi-sports academy backend.
// @host         localhost:8080
package main

import (
	"spodemy-backend/config"
	"spodemy-backend/database"
	_ "spodemy-backend/docs" // ← the generated Swagger docs
	"spodemy-backend/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)


func main() {
    // Load config
    _,_ = config.LoadConfig("config/local.json")

    // Initialize DB connection
    database.Connect()

    // Create Gin router
    r := gin.Default()
    
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Setup routes with database.DB
    routes.SetupRoutes(r, database.DB)

    r.Run(":8080")
}