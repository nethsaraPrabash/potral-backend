package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nethsaraPrabash/potral-backend/src/config"
	"github.com/nethsaraPrabash/potral-backend/src/routes"
)

func main() {
	loadEnv()
	r := gin.Default()

	db := config.ConnectDB()

	routes.Routes(r, db)

	log.Println("Server running on port 8080")
	log.Println(db)
	r.Run(":8080")
}

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("env not found")
		err = godotenv.Load()
		if err != nil {
			log.Fatal("error loading env")
		}
	}
}