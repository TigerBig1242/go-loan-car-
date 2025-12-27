package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host          = "localhost"
	port          = 5432
	username      = "user"
	password      = "password"
	database_name = "go-loan-car-db"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Connect to the database
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, database_name)

	// New logger for detailed SQL logging
	Logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	// Open a connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: Logger,
	})
	if err != nil {
		// log.Fatal(err)
		log.Printf("Database connection failed: %v", err)
		panic("Failed to connect database !")
	}

	DB = db

	fmt.Println(DB)
	fmt.Println("Database Migrated Successfully !")
}
