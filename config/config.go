package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// Config structure
type Config struct {
	MongoURI   string
	MongoDBName    string
	ServerPort string
	JWTSecret  string
}

// AppConfig is a global variable to access config values
var AppConfig Config

// LoadConfig loads environment variables from .env
func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: No .env file found. Using system environment variables.")
	}

	// Assign values to AppConfig
	AppConfig = Config{
		MongoURI:   getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDBName:    getEnv("MONGO_DB_NAME", "fiber-press"),
		ServerPort: getEnv("SERVER_PORT", "5000"),
		JWTSecret:  getEnv("JWT_SECRET", "sharif_secret_key"),
	}
}

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}


// Validator instance
var Validate = validator.New()

// ValidationError stores error details
type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

// ValidateStruct checks struct fields based on tags
func ValidateStruct(s interface{}) []*ValidationError {
	var errors []*ValidationError
	err := Validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Param(),
			}
			errors = append(errors, &element)
		}
	}
	return errors
}