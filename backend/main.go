package main

import (
	"ims-intro/models"
	"ims-intro/routers"
	"log"
	"net/http"
)

func main() {
	dataSourceName := "user=postgres port=5434 password=password dbname=ims sslmode=disable"
	models.InitDB(dataSourceName)

	router := routers.SetupRouter()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
