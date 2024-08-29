package routers

import (
	"github.com/gorilla/mux"
	"ims-intro/handlers"
	"ims-intro/middleware"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/signup", handlers.Signup).Methods("POST")

	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")

	router.HandleFunc("/products", middleware.AuthAdmin(handlers.CreateProduct)).Methods("POST")
	router.HandleFunc("/products/{id}", middleware.AuthAdmin(handlers.UpdateProduct)).Methods("PUT")
	router.HandleFunc("/products/{id}", middleware.AuthAdmin(handlers.DeleteProduct)).Methods("DELETE")

	return router
}
