package dbutil

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func ConnectDB() *gorm.DB {
	once.Do(func() {
		err := godotenv.Load(".env") // Load environment variables from .env file
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		// Read environment variables
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbIP := os.Getenv("DB_IP")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		// Construct DSN
		dsn := dbUser + ":" + dbPass + "@tcp(" + dbIP + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		instance = db
		log.Println("Connected to the database")
	})

	return instance
}
