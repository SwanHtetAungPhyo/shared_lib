package database

import (
	"fmt"
	"github.com/ProjectSMAA/commons/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=admin password=secret dbname=mydatabase port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log all SQL queries
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Get the underlying database connection for pooling configurations
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Configure connection pooling
	sqlDB.SetMaxOpenConns(20)                  // Max open connections
	sqlDB.SetMaxIdleConns(10)                  // Max idle connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Max lifetime of a connection
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)  // Max idle time of a connection

	DB = db
	fmt.Println("Connected to PostgreSQL successfully!")
}

func Migration() {
	err := DB.AutoMigrate(models.User{})
	if err != nil {
		return
	}
}
