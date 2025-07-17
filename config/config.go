package config

import (
	// "database/sql" // Old SQL package - replaced with GORM
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	// _ "github.com/lib/pq" // Old postgres driver - replaced with GORM driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB *sql.DB // Old SQL database instance - replaced with GORM
var DB *gorm.DB // GORM database instance

func ConnectDB() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get database configuration from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "mycarsdb")
	dbSSLMode := getEnv("DB_SSL_MODE", "disable")

	// Check if password is provided
	if dbPassword == "" {
		log.Fatal("DB_PASSWORD environment variable is required")
	}

	// Build connection string for GORM
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
		dbUser, dbName, dbPassword, dbHost, dbPort, dbSSLMode)

	// OLD SQL CONNECTION - replaced with GORM
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	//     log.Fatalf("Failed to connect to database: %v", err)
	// }
	// if err := db.Ping(); err != nil {
	//     log.Fatalf("Failed to ping database: %v", err)
	// }

	// NEW GORM CONNECTION
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database with GORM: %v", err)
	}

	// Test the connection using GORM
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying SQL DB: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database with GORM")
	DB = db
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
