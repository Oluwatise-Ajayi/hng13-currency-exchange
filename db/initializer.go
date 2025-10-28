// db/initializer.go
package db

import (
	// We don't need "log" here anymore
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Try to load a .env file.
	// If it doesn't exist (like on Railway), that's fine.
	// The app will just use the *real* environment variables
	// that Railway provides. We can safely ignore any errors here.
	godotenv.Load()
}