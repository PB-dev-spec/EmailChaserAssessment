package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	// Check if the .env file was loaded successfully.
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Check that all required environment variables are set.
	requiredEnvVars := []string{
		"DB_URL",
		"SECRET_KEY",
		// Add more required environment variables here as needed.
	}
	for _, envVar := range requiredEnvVars {
		if _, exists := os.LookupEnv(envVar); !exists {
			log.Fatalf("Required environment variable not set: %s", envVar)
		}
	}
}
