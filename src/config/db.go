package config

import (
	"fmt"
	"log"
	"os"

	"github.com/nethsaraPrabash/potral-backend/src/model"


	"gorm.io/driver/mysql"
	"gorm.io/gorm"	
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("✅ Database connected & migrated successfully!")
	return db
}
