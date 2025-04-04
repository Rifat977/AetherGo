package db

import (
	"AetherGo/internal/log"
	"os"

	// "path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "aetherGo.db"
	}

	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Success("Database connection established")
}

// AutoMigrate runs the migration for all detected models.
func AutoMigrate(models ...interface{}) {
	for _, model := range models {
		if !DB.Migrator().HasTable(model) {
			log.Infof("Migrating table for model: %T", model)
			err := DB.AutoMigrate(model)
			if err != nil {
				log.Fatalf("Failed to migrate model %T: %v", model, err)
			} else {
				log.Successf("Migration successful for model: %T", model)
			}
		} else {
			log.Infof("Model %T already migrated, skipping.", model)
		}
	}
}
