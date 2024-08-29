package main

import (
	"github.com/joho/godotenv"
	"ims-intro/models"
	"ims-intro/routers"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dataSourceName := "user=" + dbUser + " port=" + dbPort + " password=" + dbPassword + " dbname=" + dbName + " sslmode=" + dbSSLMode
	models.InitDB(dataSourceName)

	e := routers.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is running on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
