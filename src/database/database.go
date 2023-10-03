package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DB_INFO = map[string]interface{}{
	"USER": os.Getenv("DB_USER"),
	"PASS": os.Getenv("DB_PASS"),
	"URL":  os.Getenv("DB_URL"),
	"PORT": os.Getenv("DB_PORT"),
	"NAME": os.Getenv("DB_NAME"),
}

func Connect() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file.")
	}

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		DB_INFO["USER"].(string),
		DB_INFO["PASS"].(string),
		DB_INFO["URL"].(string),
		DB_INFO["PORT"].(string),
		DB_INFO["NAME"].(string),
	)

	fmt.Println(dsn)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database (" + dsn + ").")
	}
	DB = db
}
