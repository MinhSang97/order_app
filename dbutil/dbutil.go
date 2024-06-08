package dbutil

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func ConnectDB() *gorm.DB {
	once.Do(func() {
		err := godotenv.Load("ae.env") // Load environment variables from ae.env file
		if err != nil {
			log.Fatal("Error loading ae.env file")
		}

		// Read environment variables
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		sslMode := os.Getenv("SSL_MODE")
		timezone := os.Getenv("TIMEZONE")
		// Construct DSN
		dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=" + sslMode + " TimeZone=" + timezone

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		instance = db
		log.Println("Connected to the database")
	})

	return instance
}
