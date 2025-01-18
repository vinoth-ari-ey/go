package db

import (
	"encoding/json"
	"fmt"
	"go-crud/models"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection.
func Connect() {
	dbParams := getDBParams()

	port, err := strconv.Atoi(dbParams.Port)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbParams.Host, port, dbParams.User, dbParams.Password, dbParams.DBName)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	if err := database.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	DB = database
	log.Println("Database connected successfully!")
}

// Get db params from env
func getDBParams() models.DBParams {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		appEnv = "localhost"
	}

	// Fetch env values from vcap services since it's deployed in SAP BTP
	if appEnv != "localhost" {
		vcapServicesStr := os.Getenv("VCAP_SERVICES")
		if vcapServicesStr == "" {
			log.Fatal("VCAP_SERVICES environment variable is not set")
		}

		var vcapServices map[string]interface{}
		err := json.Unmarshal([]byte(vcapServicesStr), &vcapServices)
		if err != nil {
			log.Fatalf("Error parsing VCAP_SERVICES: %v", err)
		}

		// Ensure vcapServices is not empty
		if vcapServices == nil {
			log.Fatal("VCAP_SERVICES is empty")
		}

		// Get PostgreSQL services
		postgresServices, ok := vcapServices["postgresql-db"].([]interface{})
		if !ok || len(postgresServices) == 0 {
			log.Fatal("No postgresql-db service found in VCAP_SERVICES")
		}

		// Get the first PostgreSQL service
		firstPostgresService, ok := postgresServices[0].(map[string]interface{})
		if !ok {
			log.Fatal("First postgresql-db service is not a valid JSON object")
		}

		// Get credentials
		credentials, ok := firstPostgresService["credentials"].(map[string]interface{})
		if !ok {
			log.Fatal("No credentials found in the first postgresql-db service")
		}

		param := models.DBParams{
			Host:     credentials["hostname"].(string),
			Port:     credentials["port"].(string),
			User:     credentials["username"].(string),
			Password: credentials["password"].(string),
			DBName:   credentials["dbname"].(string),
		}
		return param
	} else {
		// Fetch values from .env
		param := models.DBParams{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		}
		return param
	}
}
