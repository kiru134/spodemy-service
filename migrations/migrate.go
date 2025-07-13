package main

import (
	"fmt"
	"log"
	"spodemy-backend/config"
	"spodemy-backend/models"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
  // 1. Load config
  cfg, err := config.LoadConfig("config/local.json")
  if err != nil {
    log.Fatalf("failed to load config: %v", err)
  }

  // 2. Connect to Postgres
  dsn := fmt.Sprintf(
    "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
    cfg.DB.Host, cfg.DB.Port,
    cfg.DB.User, cfg.DB.Password,
    cfg.DB.DBName, cfg.DB.SSLMode,
  )
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatalf("could not connect to db: %v", err)
  }

  // 3. Define versioned migrations
  migrations := []*gormigrate.Migration{
    {
      ID: "20250712_init_users_roles",
      Migrate: func(tx *gorm.DB) error {
        return tx.AutoMigrate(
          &models.Role{},
          &models.User{},
        )
      },
      Rollback: func(tx *gorm.DB) error {
        return tx.Migrator().DropTable(
          "user_roles", "users", "roles",
        )
      },
    },
    {
      ID: "20250713_create_plans_offers",
      Migrate: func(tx *gorm.DB) error {
        return tx.AutoMigrate(
          &models.Plan{},
          &models.Offer{},
        )
      },
      Rollback: func(tx *gorm.DB) error {
        return tx.Migrator().DropTable(
          "plan_offers", "offers", "plans",
        )
      },
    },
    {
      ID: "20250714_create_venues_batches",
      Migrate: func(tx *gorm.DB) error {
        return tx.AutoMigrate(
          &models.Venue{},
          &models.Batch{},
        )
      },
      Rollback: func(tx *gorm.DB) error {
        return tx.Migrator().DropTable(
          "batches", "venues",
        )
      },
    },
    {
      ID: "20250715_create_enrollments_attendance",
      Migrate: func(tx *gorm.DB) error {
        return tx.AutoMigrate(
          &models.Enrollment{},
          &models.Attendance{},
        )
      },
      Rollback: func(tx *gorm.DB) error {
        return tx.Migrator().DropTable(
          "attendance", "enrollments",
        )
      },
    },
    {
      ID: "20250716_create_fee_payments",
      Migrate: func(tx *gorm.DB) error {
        return tx.AutoMigrate(
          &models.FeePayment{},
        )
      },
      Rollback: func(tx *gorm.DB) error {
        return tx.Migrator().DropTable(
          "fee_payments",
        )
      },
    },
    {
      ID: "20250717_create_investments_transactions",
      Migrate: func(tx *gorm.DB) error {
        return tx.AutoMigrate(
          &models.Investment{},
          &models.InvestmentTransaction{},
        )
      },
      Rollback: func(tx *gorm.DB) error {
        return tx.Migrator().DropTable(
          "investment_transactions", "investments",
        )
      },
    },
    {
      ID: "20250718_create_learning",
      Migrate: func(tx *gorm.DB) error {
        return tx.AutoMigrate(
          &models.Course{},
          &models.Assessment{},
        )
      },
      Rollback: func(tx *gorm.DB) error {
        return tx.Migrator().DropTable(
          "certifications", "assessments", "courses",
        )
      },
    },
    {
      ID: "20250719_create_expenses",
      Migrate: func(tx *gorm.DB) error {
        return tx.AutoMigrate(
          &models.Expense{},
        )
      },
      Rollback: func(tx *gorm.DB) error {
        return tx.Migrator().DropTable(
          "expenses",
        )
      },
    },
    {
  ID: "20250720_reset_schema_to_uuid",
  Migrate: func(tx *gorm.DB) error {
    // 1. Drop everything (in reverse dependency order)
 if err := tx.Migrator().DropTable(
      &models.Expense{},
      &models.Assessment{}, &models.Course{},
      &models.InvestmentTransaction{}, &models.Investment{},
      &models.FeePayment{},
      &models.Attendance{}, &models.Enrollment{},
      &models.Batch{}, &models.Venue{},
      &models.Offer{}, &models.Plan{},
      &models.User{}, &models.Role{},
    ); err != nil {
      return err
    }

    // 2. (Re-)create all tables with your updated models
      return tx.AutoMigrate(
      &models.Role{}, &models.User{},
      &models.Plan{}, &models.Offer{},
      &models.Venue{}, &models.Batch{},
      &models.Enrollment{}, &models.Attendance{},
      &models.FeePayment{},
      &models.Investment{}, &models.InvestmentTransaction{},
      &models.Course{}, &models.Assessment{},
      &models.Expense{},
    )
  },
  Rollback: func(tx *gorm.DB) error {
    // nothing to roll back cleanly here
    return nil
  },
},

  }

  // 4. Run migrations
  m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
  if err := m.Migrate(); err != nil {
    log.Fatalf("migration failed: %v", err)
  }

  log.Println("Migrations applied successfully")
}