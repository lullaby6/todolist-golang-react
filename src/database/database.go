package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("[ERROR] Problem loading .env file (%s).", err)
	}

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_SERVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("[ERROR] Failed to connect database (%s) (%s).", dsn, err))
	}
	DB = db

	fmt.Printf("Connected to database sussfully! (%s)\n", dsn)
}
