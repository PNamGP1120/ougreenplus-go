package database

import (
	"log"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	var err error

	DB, err = gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	log.Println("✅ Database connected successfully")
}
