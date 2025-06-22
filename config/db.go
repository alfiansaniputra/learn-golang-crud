package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error

	// Load environment variables
	dsn := os.Getenv("DATABASE_MAIN_USER") + ":" + os.Getenv("DATABASE_MAIN_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_MAIN_HOST") + ":" + os.Getenv("DATABASE_MAIN_PORT") + ")/" + os.Getenv("DATABASE_MAIN_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connected successfully")
	return DB
}
