package main

import (
	"ims-intro/models"
	"ims-intro/routers"
	"log"
)

func main() {
	dataSourceName := "user=postgres port=5434 password=password dbname=ims sslmode=disable"
	models.InitDB(dataSourceName)

	e := routers.SetupRouter()

	log.Println("Server is running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
