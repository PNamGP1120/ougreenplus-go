package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBDSN     string
	JWTSecret string
	AWSRegion string
	S3Bucket  string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		Port:      get("PORT", "8080"),
		DBDSN:     get("DB_DSN", ""),
		JWTSecret: get("JWT_SECRET", "super_secret"),
		AWSRegion: get("AWS_REGION", "ap-southeast-1"),
		S3Bucket:  get("S3_BUCKET", ""),
	}

	if cfg.DBDSN == "" {
		log.Fatal("⚠️ Missing DB_DSN in .env")
	}
	if cfg.JWTSecret == "" {
		log.Fatal("⚠️ Missing JWT_SECRET in .env")
	}

	return cfg
}

func get(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}
