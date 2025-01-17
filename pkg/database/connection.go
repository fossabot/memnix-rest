package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	// DBConn is a pointer to gorm.DB
	DBConn   *gorm.DB
	Cache    *cache.Cache
	user     string
	password string
	host     string
	db       string
	port     string
	rabbitMQ string
)

func LoadVar() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "development" {
		log.Println("Running in development mode")
		user = os.Getenv("DEBUG_DB_USER")         // Get DB_USER from env
		password = os.Getenv("DEBUG_DB_PASSWORD") // Get DB_PASSWORD from env
		host = os.Getenv("DEBUG_DB_HOST")         // Get DB_HOST from env
		db = os.Getenv("DEBUG_DB_DB")             // Get DB_DB (db name) from env
		port = os.Getenv("DEBUG_DB_PORT")         // Get DB_PORT from env
		rabbitMQ = os.Getenv("DEBUG_RABBIT_MQ")   // Get DB_PORT from env
	} else {
		log.Println("Running in production mode")
		user = os.Getenv("DB_USER")         // Get DB_USER from env
		password = os.Getenv("DB_PASSWORD") // Get DB_PASSWORD from env
		host = os.Getenv("DB_HOST")         // Get DB_HOST from env
		db = os.Getenv("DB_DB")             // Get DB_DB (db name) from env
		port = os.Getenv("DB_PORT")         // Get DB_PORT from env
		rabbitMQ = os.Getenv("RABBIT_MQ")   // Get DB_PORT from env
	}
}

func CreateCache() error {
	Cache = cache.New(10*time.Minute, 15*time.Minute)
	return nil
}

// Connect creates a connection to database
func Connect() (err error) {
	LoadVar() // Load var from .env file

	// Convert port
	dbPort, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
		},
	)

	// Create postgres connection string
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", user, password, host, db, dbPort)
	// Open connection
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}

	sqlDB, err := DBConn.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
