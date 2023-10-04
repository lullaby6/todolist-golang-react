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
		panic(err)
	}

	DB_INFO := map[string]interface{}{
		"USER":   os.Getenv("DB_USER"),
		"PASS":   os.Getenv("DB_PASS"),
		"SERVER": os.Getenv("DB_SERVER"),
		"PORT":   os.Getenv("DB_PORT"),
		"NAME":   os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("DESKTOP-LK9AQTD://%s:%s@%s:%s?database=%s",
		DB_INFO["USER"].(string),
		DB_INFO["PASS"].(string),
		DB_INFO["SERVER"].(string),
		DB_INFO["PORT"].(string),
		DB_INFO["NAME"].(string),
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database (" + dsn + ").")
	}
	DB = db

	fmt.Printf("Connected to database sussfully! (%s)\n", dsn)
}
