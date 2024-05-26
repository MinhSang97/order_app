package dbutil

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func ConnectDB() *gorm.DB {
	once.Do(func() {
		dsn := "admin:123456@tcp(127.0.0.1:3306)/order_app?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		instance = db
		log.Println("Connected to the database")
	})

	return instance
}
