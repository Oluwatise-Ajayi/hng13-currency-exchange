package main

import (
	"log" // <-- Make sure to import "log"

	"github.com/Oluwatise-Ajayi/hng13-currency-exchange/db"
	"github.com/Oluwatise-Ajayi/hng13-currency-exchange/models"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectDB() // We'll assume this logs a fatal error if it fails
}

func main() {
	log.Println("Attempting to run database migrations...")

	// Capture the error from AutoMigrate
	err := db.DB.AutoMigrate(&models.CountryInfo{})
	
	if err != nil {
		// If there's ANY error, print it and CRASH the program.
		log.Fatalf("FATAL: Failed to run migrations: %v\n", err)
	}

	// This line will ONLY be printed if the migration was successful
	log.Println("Database migrated successfully.")
}