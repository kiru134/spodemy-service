package database

import (
	"fmt"
	"log"
	"spodemy-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection pool
var DB *gorm.DB

// Connect initializes the GORM DB connection using config settings
func Connect() {
    cfg, err := config.LoadConfig("config/local.json")
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }
    dsn := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        cfg.DB.Host, cfg.DB.Port,
        cfg.DB.User, cfg.DB.Password,
        cfg.DB.DBName, cfg.DB.SSLMode,
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    DB = db
}